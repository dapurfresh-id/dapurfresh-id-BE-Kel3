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
	// Update(ctx context.Context, user *request.RequestImgUpdate) (*entities.User, error)
	GetUser(userID string) *entities.User
	UploadImage(ctx context.Context, user *request.RequestImgUpdate) (string, error)
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

// func (service *userService) Update(ctx context.Context, user *request.RequestImgUpdate) (*entities.User, error) {
// 	upload, err := service.Upload(ctx, &request.RequestImgUpdate{Image: user.Image})
// 	if err != nil {
// 		return nil, err
// 	}
// 	userUpdate := &entities.User{
// 		ID:        user.ID,
// 		Image:     upload,
// 		UpdatedAt: time.Now(),
// 	}
// 	ctx, cancel := context.WithTimeout(ctx, service.contextTimeOut)
// 	defer cancel()
// 	res, err := service.userRepository.Update(ctx, userUpdate)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return res, nil
// }

func (service *userService) GetUser(userID string) *entities.User {
	// log.Println("tes-ser", userID)
	return service.userRepository.GetUser(userID)
}

func (service *userService) UploadImage(ctx context.Context, user *request.RequestImgUpdate) (string, error) {
	uploadFile, err := helpers.ImageUploadHelper(user.Image)
	if err != nil {
		return "", err
	}
	return uploadFile, err
}
