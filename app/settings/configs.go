package settings

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// AppSettings holds all the application settings
type AppSettings struct {
	ApiVersion            string
	ApiPath               string
	ApiDoc                string
	Host                  string
	HostGateway           string
	Port                  string
	SecretKey             string
	UsernameGateWay       string
	PasswordGateWay       string
	SecretKeyGateWay      string
	KafkaBootstrapServers string
	RedisClusterNodes     []string
	RedisHost             string
	RedisPort             string
	RedisPassword         string
}

// LoadEnv loads environment variables
func LoadEnv(envPath string) AppSettings {
	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	return AppSettings{
		ApiVersion:            os.Getenv("API_VERSION"),
		ApiPath:               os.Getenv("API_PATH"),
		ApiDoc:                os.Getenv("API_DOC"),
		Host:                  os.Getenv("HOST"),
		HostGateway:           os.Getenv("HOST_GATEWAY"),
		Port:                  os.Getenv("PORT_GINAPI_GATEWAY"),
		SecretKey:             os.Getenv("SECRET_KEY"),
		UsernameGateWay:       os.Getenv("USERNAME_GATEWAY"),
		PasswordGateWay:       os.Getenv("PASSWORD_GATEWAY"),
		SecretKeyGateWay:      os.Getenv("SECRET_KEY_GATEWAY"),
		KafkaBootstrapServers: os.Getenv("KAFKA_BOOTSTRAP_SERVERS"),
		RedisHost:             os.Getenv("HOST_REDIS"),
		RedisPort:             os.Getenv("PORT_REDIS"),
		RedisPassword:         os.Getenv("REDIS_PASSWORD"),
	}
}
