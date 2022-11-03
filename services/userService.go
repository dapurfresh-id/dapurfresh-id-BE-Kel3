package services

import (
	"log"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"github.com/aldisaputra17/dapur-fresh-id/repositories"
	"github.com/aldisaputra17/dapur-fresh-id/request"
	"github.com/mashingan/smapping"
)

type UserService interface {
	UpdateUser(user request.RequestUpdateUser) entities.User
	GetUser(userID string) *entities.User
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) UpdateUser(user request.RequestUpdateUser) entities.User {
	log.Println("tes up-ser", user)
	userToUpdate := entities.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) GetUser(userID string) *entities.User {
	log.Println("tes-ser", userID)
	return service.userRepository.GetUser(userID)
}
