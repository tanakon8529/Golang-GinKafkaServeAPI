package middleware

import (
	"context"
	"encoding/json"
	"ginapi-gateway/models"
	"ginapi-gateway/services"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Assuming your KafkaConsumer struct and NewKafkaConsumer function
// are updated to handle a slice of topics rather than a single one.

// Jobs represents a collection of background jobs.
type Jobs struct {
	KafkaConsumer *services.KafkaConsumer
}

// ReadKafkaMessages continuously reads messages from the subscribed Kafka topics.
func (j *Jobs) ReadKafkaMessages(ctx context.Context) {
	// Assuming your KafkaConsumer.Consume method is updated to handle multiple topics.
	err := j.KafkaConsumer.Consumer.SubscribeTopics(models.KafkaTopics, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topics: %v", err)
	}

	for {
		select {
		case <-ctx.Done():
			log.Println("Stopping Kafka message consumer...")
			return
		default:
			msg, err := j.KafkaConsumer.Consumer.ReadMessage(-1)
			if err != nil {
				log.Printf("Error consuming Kafka message: %v\n", err)
				continue
			}
			topic := *msg.TopicPartition.Topic
			j.handleMessage(msg.Value, topic)
			log.Printf("Received message from %s: %s\n", topic, string(msg.Value))
		}
	}
}

// handleMessage processes a single Kafka message based on its topic.
func (j *Jobs) handleMessage(msg []byte, topic string) {
	// Process the message based on the topic
	var message map[string]interface{}
	if err := json.Unmarshal(msg, &message); err != nil {
		log.Printf("Error unmarshalling JSON for topic %s: %v\n", topic, err)
		return
	}
	// Example processing
	log.Printf("Processed message from %s: %+v\n", topic, message)
}

// StartJobs initializes and starts the background jobs for Kafka message consumption.
func StartJobs() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Setting up signal capturing for graceful shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Create a Kafka consumer for multiple topics
	kafkaConsumer, err := services.NewKafkaConsumer(models.KafkaTopics, "system") // Adjust group ID as needed
	if err != nil {
		log.Fatalf("Error creating Kafka consumer: %v", err)
	}
	defer kafkaConsumer.Close()

	// Create and start Jobs
	jobs := &Jobs{KafkaConsumer: kafkaConsumer}
	go jobs.ReadKafkaMessages(ctx)

	// Wait for termination signal
	<-signals
	log.Println("Received shutdown signal, terminating...")
}
