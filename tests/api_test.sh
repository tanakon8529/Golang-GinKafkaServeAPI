
#!/bin/bash

# Set your API host
API_HOST="http://localhost:8080"

# Replace these with your actual client ID and client secret
CLIENT_ID="YourClientID"
CLIENT_SECRET="YourClientSecret"

# Auth endpoint to get the token
echo "Authenticating..."
AUTH_RESPONSE=$(curl -s -X POST "$API_HOST/auth" -H "client-id: $CLIENT_ID" -H "client-secret: $CLIENT_SECRET")
TOKEN=$(echo $AUTH_RESPONSE | jq -r '.token')
echo "Authentication token: $TOKEN"

# Health check endpoint
echo "Checking API health..."
curl -X GET "$API_HOST/health" -H "token: $TOKEN"

# Kafka endpoint to send a message
echo "Sending a message to Kafka..."
curl -X POST "$API_HOST/kafka" -H "Content-Type: application/json" -H "token: $TOKEN" -d '{"message": "Hello, Kafka!", "topic": "YourTopic"}'

echo "Tests completed."
