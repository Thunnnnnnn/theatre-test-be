package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	kafkago "github.com/segmentio/kafka-go"
)

var Writer = &kafkago.Writer{
	Addr:  kafkago.TCP("localhost:9092"),
	Topic: "notification-topic",
}

type BookingEvent struct {
	// UserID string `json:"userId"`
	Event  string `json:"event"`
	SeatID string `json:"seatId"`
}

type ReleaseEvent struct {
	Event   string   `json:"event"`
	SeatIDS []string `json:"seatIds"`
}

func PublishBookingSuccess(event BookingEvent) error {

	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = Writer.WriteMessages(
		context.Background(),
		kafkago.Message{
			Value: data,
		},
	)

	if err != nil {
		fmt.Println("Kafka Publish Error:", err)
		return err
	}

	fmt.Println("Kafka Event Published")

	return nil
}

func PublishReleaseEvent(event ReleaseEvent) error {
	fmt.Println("Publishing Release Event:", event)
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = Writer.WriteMessages(
		context.Background(),
		kafkago.Message{
			Value: data,
		},
	)

	if err != nil {
		fmt.Println("Kafka Publish Error:", err)
		return err
	}

	fmt.Println("Kafka Event Published")

	return nil
}
