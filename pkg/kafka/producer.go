package kafka

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer struct {
	ConfigMap *ckafka.ConfigMap
}

func NewKafkaProducer(configMap *ckafka.ConfigMap) *Producer {
	return &Producer{ConfigMap: configMap}
}

func (p *Producer) Publish(msg interface{}, key []byte, topic string, deliveryChan chan kafka.Event) error {
	producer, err := ckafka.NewProducer(p.ConfigMap)
	if err != nil {
		return err
	}

	msgJson, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny},
		Value:          msgJson,
		Key:            key,
	}
	err = producer.Produce(message, deliveryChan)
	if err != nil {
		panic(err)
	}

	producer.Flush(1000)
	return nil
}
