package config

import (
	"os"

	"github.com/joho/godotenv"
)

func CheckLoadEnvEmail() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed load env file")
	}
}

func SmptHost() string {
	CheckLoadEnvEmail()
	return os.Getenv("CONFIG_SMTP_HOST")
}

func SmptPort() string {
	CheckLoadEnvEmail()
	return os.Getenv("CONFIG_SMTP_PORT")
}

func SenderName() string {
	CheckLoadEnvEmail()
	return os.Getenv("CONFIG_SENDER_NAME")
}

func AuthEmail() string {
	CheckLoadEnvEmail()
	return os.Getenv("CONFIG_AUTH_EMAIL")
}

func AuthPassword() string {
	CheckLoadEnvEmail()
	return os.Getenv("CONFIG_AUTH_PASSWORD")
}
