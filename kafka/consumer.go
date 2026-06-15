package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	ws "theatre-test-api/websocket"

	"github.com/gorilla/websocket"
	kafkago "github.com/segmentio/kafka-go"
)

type HoldedExpireEvent struct {
	Event   string   `json:"event"`
	SeatID  string   `json:"seatId"`
	SeatIDS []string `json:"seatIds"`
}

func StartConsumer() {

	reader := kafkago.NewReader(
		kafkago.ReaderConfig{
			Brokers: []string{"localhost:9092"},
			Topic:   "notification-topic",
			GroupID: "notification-group",
		},
	)

	for {

		msg, err := reader.ReadMessage(
			context.Background(),
		)

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Event:", string(msg.Value))

		var payload HoldedExpireEvent

		err = json.Unmarshal(msg.Value, &payload)

		fmt.Println("Payload:", payload)

		if err != nil {
			log.Println("invalid kafka message:", err)
			return
		}

		data, _ := json.Marshal(map[string]interface{}{
			"event":   payload.Event,
			"seatId":  payload.SeatID,
			"seatIds": payload.SeatIDS,
		})

		for client := range ws.WSHub.Clients {

			err := client.WriteMessage(
				websocket.TextMessage,
				data,
			)

			if err != nil {
				client.Close()
				delete(ws.WSHub.Clients, client)
			}
		}
	}
}
