package services

import (
	"context"
	"time"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"github.com/aldisaputra17/dapur-fresh-id/repositories"
	"github.com/aldisaputra17/dapur-fresh-id/request"
)

type UserService interface {
	Update(ctx context.Context, user *request.RequestUserUpdate) (*entities.User, error)
	GetUser(userID string) *entities.User
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

func (service *userService) Update(ctx context.Context, user *request.RequestUserUpdate) (*entities.User, error) {
	userUpdate := &entities.User{
		ID:        user.ID,
		Username:  user.Username,
		Name:      user.Name,
		Phone:     user.Phone,
		Password:  user.Password,
		ImageID:   user.ImageID,
		UpdatedAt: time.Now(),
	}
	ctx, cancel := context.WithTimeout(ctx, service.contextTimeOut)
	defer cancel()
	res, err := service.userRepository.Update(ctx, userUpdate)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (service *userService) GetUser(userID string) *entities.User {
	// log.Println("tes-ser", userID)
	return service.userRepository.GetUser(userID)
}
