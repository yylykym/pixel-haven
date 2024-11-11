package event

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"pixel-haven/interval/config"
)

func Consumer(topic string) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{config.Kafka_Broker},
		Topic:    topic,
		GroupID:  config.Group_Id,
		MinBytes: 1,    // 10KB
		MaxBytes: 10e6, // 10MB
	})
	defer func() {
		err := reader.Close()
		if err != nil {
			log.Fatalf("Error closing reader: %s\n", err)
		}
	}()
	for {
		// 读取消息
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("could not read message %v", err)
			break
		}

		fmt.Printf("consumed: %s = %s\n", string(msg.Key), string(msg.Value))
		// 消费者确认消息已处理
		err = reader.CommitMessages(context.Background(), msg)
		if err != nil {
			log.Printf("Failed to commit message: %v", err)
		}
	}
}
