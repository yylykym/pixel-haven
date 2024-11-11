package event

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"pixel-haven/interval/config"
)

func Producer(topic string, key string, value string) {
	writer := &kafka.Writer{
		Addr:         kafka.TCP(config.Kafka_Broker),
		Topic:        topic,
		Balancer:     &kafka.Hash{}, // 使用Hash Balancer，可以选择适合你的分区策略
		BatchTimeout: 100,
		BatchSize:    100,
	}
	defer func() {
		err := writer.Close()
		if err != nil {
			log.Fatal("Could not close writer", err)
		}
	}()
	msg := kafka.Message{
		Key:   []byte(fmt.Sprintf("Key-%d", key)),
		Value: []byte(fmt.Sprintf("Message-%d", value)),
	}

	// 发送消息到Kafka
	err := writer.WriteMessages(context.Background(), msg)
	if err != nil {
		log.Fatalf("could not write message %v", err)
	} else {
		log.Printf("produced: %s\n", msg.Value)
	}
}
