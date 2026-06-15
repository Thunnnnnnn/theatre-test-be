package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"theatre-test-api/database"
	"theatre-test-api/kafka"
	"theatre-test-api/models"
	ws "theatre-test-api/websocket"

	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindAllSeats() ([]models.Seat, error) {
	ctx := context.Background()

	if database.SeatCollection == nil {
		return nil, nil
	}

	cursor, err := database.SeatCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var seats []models.Seat
	if err := cursor.All(ctx, &seats); err != nil {
		return nil, err
	}

	return seats, nil
}

func FindSeatById(id primitive.ObjectID) (models.Seat, error) {
	ctx := context.Background()

	if database.SeatCollection == nil {
		return models.Seat{}, nil
	}

	var seat models.Seat
	err := database.SeatCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&seat)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Seat{}, nil
		}
		return models.Seat{}, err
	}

	return seat, nil
}

func CreateSeat(seat models.Seat) (models.Seat, error) {
	ctx := context.Background()

	if database.SeatCollection == nil {
		return models.Seat{}, nil
	}

	seat.CreatedAt = time.Now()
	seat.UpdatedAt = time.Now()
	seat.IsBooked = false
	seat.IsHolded = false
	seat.Status = "available"

	result, err := database.SeatCollection.InsertOne(ctx, seat)
	if err != nil {
		return models.Seat{}, err
	}

	seat.ID = result.InsertedID.(primitive.ObjectID)
	return seat, nil
}

func UpdateSeat(id primitive.ObjectID, seat models.Seat) (models.Seat, error) {
	ctx := context.Background()

	if database.SeatCollection == nil {
		return models.Seat{}, nil
	}

	findSeat, err := FindSeatById(id)
	if err != nil {
		return models.Seat{}, err
	}

	seat.CreatedAt = findSeat.CreatedAt
	seat.UpdatedAt = time.Now()
	seat.IsBooked = findSeat.IsBooked

	_, err = database.SeatCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": seat})
	if err != nil {
		return models.Seat{}, err
	}

	return FindSeatById(id)
}

func DeleteSeat(id primitive.ObjectID) error {
	ctx := context.Background()

	if database.SeatCollection == nil {
		return nil
	}
	_, err := database.SeatCollection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func FindSeatsByTheatreId(theatreId primitive.ObjectID) ([]models.Seat, error) {
	ctx := context.Background()

	if database.SeatCollection == nil {
		return nil, nil
	}

	cursor, err := database.SeatCollection.Find(ctx, bson.M{"theatre_id": theatreId})
	if err != nil {
		return nil, err
	}

	var seats []models.Seat
	if err := cursor.All(ctx, &seats); err != nil {
		return nil, err
	}

	return seats, nil
}

func FindSeatsByIsBooked(keyword string) ([]models.AdminSeatResponse, error) {
	ctx := context.Background()
	pipeline := mongo.Pipeline{
		{
			{"$match", bson.D{
				{"is_booked", true},
			}},
		},
		{
			{"$lookup", bson.D{
				{"from", "theatres"},
				{"localField", "theatre_id"},
				{"foreignField", "_id"},
				{"as", "theatre"},
			}},
		},
		{
			{"$unwind", "$theatre"},
		},
		{
			{"$lookup", bson.D{
				{"from", "movies"},
				{"localField", "theatre.movie_id"},
				{"foreignField", "_id"},
				{"as", "movie"},
			}},
		},
		{
			{"$unwind", "$movie"},
		},
		{
			{"$match", bson.D{
				{"$or", bson.A{
					bson.D{
						{"theatre.theatre_name", bson.D{
							{"$regex", keyword},
							{"$options", "i"},
						}},
					},
					bson.D{
						{"movie.name", bson.D{
							{"$regex", keyword},
							{"$options", "i"},
						}},
					},
				}},
			}},
		},
		{
			{"$sort", bson.D{
				{"holded_expire", -1},
			}},
		},
	}

	if database.SeatCollection == nil {
		return nil, nil
	}

	cursor, err := database.SeatCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	var seats []models.Seat
	if err := cursor.All(ctx, &seats); err != nil {
		return nil, err
	}
	var seatsResponse []models.AdminSeatResponse

	for _, seat := range seats {
		theatre, err := FindTheatreById(seat.TheatreId)
		if err != nil {
			log.Println("FindTheatreById error:", err)
		}

		var user models.User
		if seat.BookedById != nil {
			user, err = FindUserById(*seat.BookedById)
			if err != nil {
				log.Println("FindUserById error:", err)
			}
		}

		seatsResponse = append(seatsResponse, models.AdminSeatResponse{
			ID:           seat.ID,
			TheatreId:    seat.TheatreId,
			Theatre:      theatre,
			Row:          seat.Row,
			Number:       seat.Number,
			Status:       seat.Status,
			IsHolded:     seat.IsHolded,
			HoldedById:   seat.HoldedById,
			HoldedExpire: seat.HoldedExpire,
			IsBooked:     seat.IsBooked,
			BookedById:   seat.BookedById,
			BookedBy:     user,
			CreatedAt:    seat.CreatedAt,
			UpdatedAt:    seat.UpdatedAt,
		})
	}

	// var seats []models.Seat
	// if err := cursor.All(ctx, &seats); err != nil {
	// 	return nil, err
	// }

	return seatsResponse, nil
}

func UpdateBookedStatus(id primitive.ObjectID, userID string) (models.Seat, error) {
	ctx := context.Background()

	if database.SeatCollection == nil {
		return models.Seat{}, nil
	}

	findSeat, err := FindSeatById(id)
	if err != nil {
		return models.Seat{}, err
	}

	if findSeat.IsBooked {
		return models.Seat{}, fmt.Errorf("seat already booked")
	}

	holdId := findSeat.HoldedById

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return models.Seat{}, err
	}

	if !findSeat.IsHolded {
		return models.Seat{}, fmt.Errorf("please hold the seat first")
	}

	if findSeat.IsHolded && holdId != nil && *holdId != userObjectID {
		return models.Seat{}, fmt.Errorf("seat is holded by another user")
	}

	if findSeat.HoldedExpire.Before(time.Now()) {
		return models.Seat{}, fmt.Errorf("hold expired")
	}

	_, err = database.SeatCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{
		"is_booked":    true,
		"booked_by_id": userID,
		"status":       "booked",
		"updated_at":   time.Now(),
	}})
	if err != nil {
		return models.Seat{}, err
	}

	event := kafka.BookingEvent{
		Event:  "booking_success",
		SeatID: id.Hex(),
	}

	err = kafka.PublishBookingSuccess(event)

	if err != nil {
		fmt.Println("Kafka Error:", err)
	} else {
		fmt.Println("Kafka Event Sent")
	}

	return FindSeatById(id)
}

func UpdateIsHoldedStatus(id primitive.ObjectID, userID string) (models.Seat, error) {
	ctx := context.Background()

	if database.SeatCollection == nil {
		return models.Seat{}, nil
	}

	findSeat, err := FindSeatById(id)
	if err != nil {
		return models.Seat{}, err
	}

	if findSeat.IsHolded {
		return models.Seat{}, fmt.Errorf("seat already holded")
	}

	_, err = database.SeatCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": bson.M{
		"is_holded":     true,
		"holded_by_id":  userID,
		"status":        "holded",
		"holded_expire": time.Now().Add(5 * time.Minute),
		"updated_at":    time.Now(),
	}})

	if err != nil {
		return models.Seat{}, err
	}

	_, err = database.TheatreCollection.UpdateOne(ctx, bson.M{"_id": findSeat.TheatreId}, bson.M{"$inc": bson.M{"available_seats": -1}})

	if err != nil {
		return models.Seat{}, err
	}

	for client := range ws.WSHub.Clients {
		msg := map[string]interface{}{
			"event":   "seat_holded",
			"seat_id": id.Hex(),
		}

		data, _ := json.Marshal(msg)

		err := client.WriteMessage(
			websocket.TextMessage,
			data,
		)

		if err != nil {
			client.Close()
			delete(ws.WSHub.Clients, client)
		}
	}

	return FindSeatById(id)
}

func GetSeatUserByIsBookedAndIsHolded(userID primitive.ObjectID) (FindSeats []models.SeatResponse, err error) {
	ctx := context.Background()

	if database.SeatCollection == nil {
		return nil, nil
	}

	cursor, err := database.SeatCollection.Find(ctx, bson.M{"is_holded": true},
		options.Find().SetSort(bson.D{
			{Key: "holded_expire", Value: -1}, // DESC
		}))
	if err != nil {
		return nil, err
	}

	var seats []models.Seat
	if err := cursor.All(ctx, &seats); err != nil {
		return nil, err
	}

	var seatsResponse []models.SeatResponse
	for _, seat := range seats {

		theatre, err := FindTheatreById(seat.TheatreId)
		if err != nil {
			log.Println("FindTheatreById error:", err)
		}
		seatsResponse = append(seatsResponse, models.SeatResponse{
			ID:           seat.ID,
			TheatreId:    seat.TheatreId,
			Theatre:      theatre,
			Row:          seat.Row,
			Number:       seat.Number,
			Status:       seat.Status,
			IsHolded:     seat.IsHolded,
			HoldedById:   seat.HoldedById,
			HoldedExpire: seat.HoldedExpire,
			IsBooked:     seat.IsBooked,
			BookedById:   seat.BookedById,
			CreatedAt:    seat.CreatedAt,
			UpdatedAt:    seat.UpdatedAt,
		})
	}

	return seatsResponse, nil
}

func UpdateReleaseHoldedSeat() (err error) {
	now := time.Now()
	ctx := context.Background()

	filter := bson.M{
		"status": "holded",
		"holded_expire": bson.M{
			"$lt": now,
		},
	}

	update := bson.M{
		"$set": bson.M{
			"status":    "available",
			"is_holded": false,
		},
		"$unset": bson.M{
			"holded_expire": "",
			"holded_by_id":  "",
		},
	}

	var expiredSeats []models.Seat

	cursor, err := database.SeatCollection.Find(ctx, filter)
	if err != nil {
		return err
	}

	if err := cursor.All(ctx, &expiredSeats); err != nil {
		return err
	}

	theatreCount := make(map[primitive.ObjectID]int)

	for _, seat := range expiredSeats {
		theatreCount[seat.TheatreId]++
	}

	TheatreCollection := database.TheatreCollection

	for theatreID, count := range theatreCount {
		_, err := TheatreCollection.UpdateOne(
			ctx,
			bson.M{
				"_id": theatreID,
			},
			bson.M{
				"$inc": bson.M{
					"available_seats": count,
				},
			},
		)

		if err != nil {
			return err
		}
	}

	result, err := database.SeatCollection.UpdateMany(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount > 0 {
		println("released seats:", result.ModifiedCount)

		event := kafka.ReleaseEvent{
			Event:   "seat_released",
			SeatIDS: make([]string, len(expiredSeats)),
		}
		for i, seat := range expiredSeats {
			event.SeatIDS[i] = seat.ID.Hex()
		}

		err = kafka.PublishReleaseEvent(event)

		if err != nil {
			fmt.Println("Kafka Error:", err)
		} else {
			fmt.Println("Kafka Event Sent")
		}
	}

	return nil
}
