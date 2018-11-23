package cnckafkaapi

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"fmt"
	"os"
	"twittersearch/config"
	"encoding/json"
)

//this function post messages to Kafka Topic. uses confluent kafka package to connect.
func PostMessageToKafka (msg config.Single_Tweet_result, topic string){

	newmsg, err := json.Marshal(msg)
	kproducer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		fmt.Println("unable to connect to Kafka. Check Config.")
		os.Exit(-1)
	}

	defer kproducer.Close()

	kproducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(newmsg),
	}, nil)
	// Wait for message deliveries before shutting down
	kproducer.Flush(1 * 100)
}
