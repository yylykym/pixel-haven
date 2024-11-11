package event

import (
	"github.com/segmentio/kafka-go"
	"net"
	"pixel-haven/interval/config"
	"strconv"
)

func CreateTopic(topic string) error {
	conn, err := kafka.Dial("tcp", config.Kafka_Broker)
	if err != nil {
		return err
	}
	defer conn.Close()
	controller, err := conn.Controller()
	if err != nil {
		return err
	}
	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}
	return nil
}

func StartConsumer() {
	Consumer("upload")
}
