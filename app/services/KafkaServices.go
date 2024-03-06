package services

import (
	"encoding/json"
	"fmt"
	"log"

	"ginapi-gateway/settings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// KafkaProducer represents a Kafka message producer.
type KafkaProducer struct {
	Producer *kafka.Producer
	Topic    string
}

func NewKafkaProducer(topic string) (*KafkaProducer, error) {
	config, err := loadKafkaConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load Kafka config: %w", err)
	}

	p, err := kafka.NewProducer(config) // Pass the pointer directly
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka producer: %w", err)
	}

	return &KafkaProducer{Producer: p, Topic: topic}, nil
}

// Produce sends a message to the Kafka topic.
func (kp *KafkaProducer) Produce(record interface{}) error {
	encodedRecord, err := json.Marshal(record)
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %w", err)
	}

	return kp.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kp.Topic, Partition: kafka.PartitionAny},
		Value:          encodedRecord,
	}, nil)
}

// Flush waits for all messages in the Producer queue to be delivered.
func (kp *KafkaProducer) Flush() {
	kp.Producer.Flush(15000) // 15 seconds in milliseconds
}

// Close closes the Producer instance.
func (kp *KafkaProducer) Close() {
	kp.Producer.Close()
}

// KafkaConsumer represents a Kafka message consumer.
type KafkaConsumer struct {
	Consumer *kafka.Consumer
}

// NewKafkaConsumer creates a new KafkaConsumer instance for multiple topics.
func NewKafkaConsumer(topics []string, groupID string) (*KafkaConsumer, error) {
	config, err := loadKafkaConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load Kafka config: %w", err)
	}

	// Correctly set configuration values
	(*config)["group.id"] = groupID
	(*config)["auto.offset.reset"] = "earliest"

	c, err := kafka.NewConsumer(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka consumer: %w", err)
	}

	// Subscribe to the topics
	if err := c.SubscribeTopics(topics, nil); err != nil {
		c.Close() // Ensure resources are released on failure
		return nil, fmt.Errorf("failed to subscribe to topics: %w", err)
	}

	return &KafkaConsumer{Consumer: c}, nil
}

// Consume reads messages from the subscribed topics.
func (kc *KafkaConsumer) Consume() ([]byte, error) {
	msg, err := kc.Consumer.ReadMessage(-1) // Block until a message is received
	if err != nil {
		return nil, fmt.Errorf("consumer error: %w", err)
	}

	return json.Marshal(msg.Value)
}

// Close closes the Consumer instance.
func (kc *KafkaConsumer) Close() {
	kc.Consumer.Close()
}

// loadKafkaConfig loads the Kafka configuration from settings and returns a pointer to kafka.ConfigMap.
func loadKafkaConfig() (*kafka.ConfigMap, error) {
	config := settings.LoadEnv(".env")
	log.Printf("KafkaBootstrapServers: %s\n", config.KafkaBootstrapServers)
	// Ensure security.protocol is set to PLAINTEXT
	return &kafka.ConfigMap{
		"bootstrap.servers": config.KafkaBootstrapServers,
		"security.protocol": "PLAINTEXT", // Explicitly set to PLAINTEXT
	}, nil
}
