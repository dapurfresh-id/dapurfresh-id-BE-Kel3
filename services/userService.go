package services

import (
	"context"
	"time"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"github.com/aldisaputra17/dapur-fresh-id/helpers"
	"github.com/aldisaputra17/dapur-fresh-id/repositories"
	"github.com/aldisaputra17/dapur-fresh-id/request"
)

type UserService interface {
	Update(ctx context.Context, user *request.RequestUserUpdate, img interface{}) (*entities.User, error)
	File(file *request.RequestUserUpdate) (string, error)
	RemoteUpload(image *entities.User) (string, error)
}

type userService struct {
	userRepository repositories.UserRepository
	contextTimeOut time.Duration
}

func NewUserService(userRepo repositories.UserRepository, time time.Duration) UserService {
	return &userService{
		userRepository: userRepo,
		contextTimeOut: time,
	}
}

func (service *userService) File(file *request.RequestUserUpdate) (string, error) {
	uploadUrl, err := helpers.ImageUploadHelper(file.Image)
	if err != nil {
		return "", err
	}

	return uploadUrl, nil
}

func (service *userService) RemoteUpload(image *entities.User) (string, error) {
	uploadUrl, errUrl := helpers.ImageUploadHelper(image.Image)
	if errUrl != nil {
		return "", errUrl
	}

	return uploadUrl, nil
}

func (service *userService) Update(ctx context.Context, user *request.RequestUserUpdate, img interface{}) (*entities.User, error) {
	userUpdate := &entities.User{
		ID:        user.ID,
		Username:  user.Username,
		Name:      user.Name,
		Phone:     user.Phone,
		Password:  user.Password,
		Image:     user.Image.Filename,
		UpdatedAt: time.Now(),
	}
	ctx, cancel := context.WithTimeout(ctx, service.contextTimeOut)
	defer cancel()

	upload, err := helpers.ImageUploadHelper(user.Image)
	if err != nil {
		return nil, err
	}
	res, err := service.userRepository.Update(ctx, userUpdate, upload)
	if err != nil {
		return nil, err
	}
	return res, nil
}
