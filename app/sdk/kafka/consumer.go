package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer(brokers []string, topic, groupId string) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  groupId,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	return &Consumer{
		reader: r,
	}
}

func (c *Consumer) ReadMessage(ctx context.Context, handler func(msg kafka.Message)) error {
	for {
		m, err := c.reader.ReadMessage(ctx)
		if err != nil {
			return err
		}
		handler(m)
	}
}

func (p *Consumer) Close() {
	if err := p.reader.Close(); err != nil {
		log.Printf("failed to close kafka writer: %v", err)
	}
}
