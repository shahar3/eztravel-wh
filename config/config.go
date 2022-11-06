package config

import (
	"github.com/joho/godotenv"
	"os"
)

func EnvMongoURI() string {
	_ = godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}

	return os.Getenv("MONGOURI")
}
