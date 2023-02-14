package kafkapackage

import (
	"log"
	"os"
)
kafka

import (
"log"
"os"

ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)


// NewKafkaProducer cria uma instância kafka.Producer pronta para uso

func NewKafkaProducer() *ckafka.Producer {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		// "security.protocol": os.Getenv("security.protocol"),
		// "sasl.mechanisms":   os.Getenv("sasl.mechanisms"),
		// "sasl.username":     os.Getenv("sasl.username"),
		// "sasl.password":     os.Getenv("sasl.password"),
	}
	p, err := ckafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	return p
}


// Publish é uma função simples criada para publicar uma nova mensagem no kafka

func Publish(msg string, topic string, producer *ckafka.Producer) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          []byte(msg),
	}
	err := producer.Produce(message, nil)
	if err != nil {
		return err
	}
	return nil
}

