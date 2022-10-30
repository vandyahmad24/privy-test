package util

import (
	"encoding/json"
	"os"

	"github.com/joho/godotenv"
)

func Stringify(data interface{}) string {
	dataByte, _ := json.Marshal(data)
	return string(dataByte)
}

func GetEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	return os.Getenv(key)
}
