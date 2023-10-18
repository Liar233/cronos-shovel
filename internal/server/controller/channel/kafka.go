package channel

import (
	"context"

	"github.com/Liar233/cronos-shovel/internal/server/model"
	"github.com/segmentio/kafka-go"
)

type KafkaChannelConfig struct {
	Brokers   []string `yaml:"brokers"`
	Topic     string   `yaml:"topic"`
	Partition int      `yaml:"partition"`
}

type KafkaChannel struct {
	writer    *kafka.Writer
	topic     string
	partition int
}

func NewKafkaChannel(config KafkaChannelConfig) ChannelInterface {

	writer := &kafka.Writer{
		Addr:                   kafka.TCP(config.Brokers...),
		Topic:                  config.Topic,
		AllowAutoTopicCreation: true,
	}

	return &KafkaChannel{
		writer:    writer,
		topic:     config.Topic,
		partition: config.Partition,
	}
}

func (kc KafkaChannel) Fire(ctx context.Context, messages []*model.Message) error {

	bulk := make([]kafka.Message, len(messages))

	for i, msg := range messages {

		bulk[i] = kafka.Message{
			Topic:     kc.topic,
			Partition: kc.partition,
			Key:       nil,
			Value:     msg.Payload,
		}
	}

	return kc.writer.WriteMessages(ctx, bulk...)
}

func (kc KafkaChannel) Close() error {

	return kc.writer.Close()
}
