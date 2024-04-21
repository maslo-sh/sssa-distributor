package broker

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"privileges-management/config"
	"strconv"
)

func CreateKafkaWriter(topicSuffix string) *kafka.Writer {
	conf := config.GetConfig()
	return &kafka.Writer{
		Addr:     kafka.TCP(fmt.Sprintf("%s:%s", conf.Kafka.Host, strconv.Itoa(conf.Kafka.Port))),
		Topic:    conf.Kafka.WriterTopic + "_" + topicSuffix,
		Balancer: &kafka.LeastBytes{},
	}
}

func WriteMessage(w *kafka.Writer, key, message string) error {
	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(key),
			Value: []byte(message),
		})

	return err
}
