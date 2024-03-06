
import requests

# Set your API host
API_HOST = "http://localhost:8080"

# Replace these with your actual client ID and client secret
headers = {
    "client-id": "YourClientID",
    "client-secret": "YourClientSecret",
}

# Function to authenticate and get token
def authenticate():
    response = requests.post(f"{API_HOST}/auth", headers=headers)
    return response.json().get("token")

# Function to check API health
def check_health(token):
    headers = {"token": token}
    response = requests.get(f"{API_HOST}/health", headers=headers)
    return response.status_code

# Function to send a message to Kafka
def send_kafka_message(token):
    headers = {
        "Content-Type": "application/json",
        "token": token,
    }
    data = {
        "message": "Hello, Kafka!",
        "topic": "YourTopic",
    }
    response = requests.post(f"{API_HOST}/kafka", headers=headers, json=data)
    return response.status_code

# Main function to run the tests
def main():
    print("Authenticating...")
    token = authenticate()
    print(f"Token: {token}")

    print("Checking API health...")
    health_status = check_health(token)
    print(f"Health status code: {health_status}")

    print("Sending a message to Kafka...")
    kafka_status = send_kafka_message(token)
    print(f"Kafka status code: {kafka_status}")

if __name__ == "__main__":
    main()
