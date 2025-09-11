package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(brokers []string, topic string) *Producer {
	w := kafka.NewWriter(kafka.WriterConfig{
		Balancer: &kafka.LeastBytes{},
		Topic:    topic,
		Brokers:  brokers,
	})

	return &Producer{
		writer: w,
	}
}

func (p *Producer) WriteMessage(ctx context.Context, key, value []byte) error {
	msg := kafka.Message{
		Key:   key,
		Value: value,
	}

	return p.writer.WriteMessages(ctx, msg)
}

func (p *Producer) SendMessages(ctx context.Context, msgs []kafka.Message) error {
	return p.writer.WriteMessages(ctx, msgs...)
}

func (p *Producer) Close() {
	if err := p.writer.Close(); err != nil {
		log.Printf("failed to close kafka writer: %v", err)
	}
}
