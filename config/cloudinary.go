package config

import (
	"os"
)

func EnvCloudName() string {
	return os.Getenv("CLOUDINARY_CLOUD_NAME")
}

func EnvCloudAPIKey() string {
	return os.Getenv("CLOUDINARY_API_KEY")
}

func EnvCloudAPISecret() string {
	return os.Getenv("CLOUDINARY_API_SECRET")
}

func EnvCloudUploadFolder() string {
	return os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
}

func EnvCloudUrl() string {
	return os.Getenv("CLOUDINARY_URL")
}
