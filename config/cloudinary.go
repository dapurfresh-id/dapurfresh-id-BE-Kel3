package config

import (
	"os"

	"github.com/joho/godotenv"
)

func CheckLoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load env file")
	}
}

func EnvCloudName() string {
	CheckLoadEnv()
	return os.Getenv("CLOUDINARY_CLOUD_NAME")
}

func EnvCloudAPIKey() string {
	CheckLoadEnv()
	return os.Getenv("CLOUDINARY_API_KEY")
}

func EnvCloudAPISecret() string {
	CheckLoadEnv()
	return os.Getenv("CLOUDINARY_API_SECRET")
}

func EnvCloudUploadFolder() string {
	CheckLoadEnv()
	return os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
}
