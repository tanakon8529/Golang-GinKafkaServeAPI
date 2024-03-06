package controllers

import (
	"ginapi-gateway/models"
	"ginapi-gateway/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SendKafkaMessage sends a message to a Kafka topic.
// @Summary Send message to Kafka
// @Description Send a message to a Kafka topic
// @Tags kafka
// @Accept  json
// @Produce  json
// @Param  Authorization  header  string  true  "token"
// @Param  KafkaRequest body models.KafkaRequest true "topic message"
// @Success 200 {object} map[string]string
// @Router /kafka [post]
func SendKafkaMessage(c *gin.Context) {
	var kafkaRequest models.KafkaRequest

	if err := c.ShouldBindJSON(&kafkaRequest); err != nil {
		respondWithError(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	if !models.Contains(models.KafkaTopics, kafkaRequest.Topic) {
		respondWithError(c, http.StatusBadRequest, "Invalid topic")
		return
	}

	kafkaProducer, err := services.NewKafkaProducer(kafkaRequest.Topic)
	if err != nil {
		respondWithError(c, http.StatusInternalServerError, "Failed to create Kafka producer")
		return
	}
	defer kafkaProducer.Close() // Ensure the producer is closed when the function returns

	if err := kafkaProducer.Produce(kafkaRequest.Message); err != nil {
		respondWithError(c, http.StatusInternalServerError, "Failed to send message to Kafka")
		return
	}

	kafkaProducer.Flush()
	c.JSON(http.StatusOK, gin.H{"message": "Message sent to Kafka"})
}

// respondWithError sends an error response.
func respondWithError(c *gin.Context, statusCode int, errorMsg string) {
	c.JSON(statusCode, gin.H{"error": errorMsg})
}
