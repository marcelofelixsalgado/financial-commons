package kafka

import (
	"testing"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/stretchr/testify/assert"
)

func TestProducerPublish(t *testing.T) {
	type TransactionDtoOutput struct {
		ID           string `json:"id"`
		Status       string `json:"status"`
		ErrorMessage string `json:"error_message"`
	}

	expectedOutput := TransactionDtoOutput{
		ID:           "1",
		Status:       "rejected",
		ErrorMessage: "you dont have limit for this transaction",
	}
	// outputJson, _ := json.Marshal(expectedOutput)

	configMap := ckafka.ConfigMap{
		"test.mock.num.brokers": 3,
	}
	producer := NewKafkaProducer(&configMap)

	deliveryChan := make(chan ckafka.Event)
	err := producer.Publish(expectedOutput, []byte("1"), "test", deliveryChan)

	e := <-deliveryChan

	msg := e.(*ckafka.Message)

	assert.Nil(t, err)
	assert.NotNil(t, e)
	assert.NotNil(t, msg)
	assert.Nil(t, msg.TopicPartition.Error)
}
