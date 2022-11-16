package services

import (
	"github.com/aldisaputra17/dapur-fresh-id/helpers"
	"github.com/aldisaputra17/dapur-fresh-id/request"
)

type ImageService interface {
	Create(file request.ImageRequest) (string, error)
}

type media struct {
}

func NewImage() ImageService {
	return &media{}
}

func (service *media) Create(file request.ImageRequest) (string, error) {
	uploadFile, err := helpers.ImageUploadHelper(file.File)
	if err != nil {
		return "", err
	}
	return uploadFile, nil
}
