// Model post go for send request to kafka
package models

import (
	"encoding/json"
)

// Topic string
// Message json
type KafkaRequest struct {
	Topic   string          `json:"topic"`
	Message json.RawMessage `json:"message"`
}

// List string array Kafka Topics can be added here
var KafkaTopics = []string{
	"daily_transaction",
	"weekly_transaction",
	"monthly_transaction",
	"yearly_transaction",
}

// Contains checks if a string is present in a slice.
func Contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
