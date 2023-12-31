package kafka

import (
	"github.com/segmentio/kafka-go"
	"legocy-go/internal/config"
)

func NewKafkaConsumer(topicName string, partition int) *kafka.Reader {
	return kafka.NewReader(
		kafka.ReaderConfig{
			Brokers:   []string{config.AppConfigInstance.KafkaConf.URI},
			Topic:     topicName,
			Partition: partition,
			MaxBytes:  10e6,
		})
}
