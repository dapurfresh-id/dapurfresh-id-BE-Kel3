package services

import (
	"context"
	"log"
	"time"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"github.com/aldisaputra17/dapur-fresh-id/repositories"
	"github.com/aldisaputra17/dapur-fresh-id/request"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(username string, password string) interface{}
	CreateUser(ctx context.Context, userReq request.RequestRegister) (*entities.User, error)
	IsDuplicateUsername(username string) bool
}

type authService struct {
	userRepository repositories.UserRepository
	contextTimeout time.Duration
}

func NewAuthService(userRepo repositories.UserRepository, time time.Duration) AuthService {
	return &authService{
		userRepository: userRepo,
		contextTimeout: time,
	}
}

func (service *authService) VerifyCredential(username string, password string) interface{} {
	res := service.userRepository.VerifyCredential(username, password)
	if v, ok := res.(entities.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Username == username && comparedPassword {
			return res
		}
		return false
	}
	return false

}

func (service *authService) CreateUser(ctx context.Context, userReq request.RequestRegister) (*entities.User, error) {
	// id, err := uuid.NewRandom()
	// if err != nil {
	// 	return nil, err
	// }
	userCreate := &entities.User{
		ID:        uuid.NewString(),
		Name:      userReq.Name,
		Username:  userReq.Username,
		Password:  userReq.Password,
		Phone:     userReq.Phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	ctx, cancel := context.WithTimeout(ctx, service.contextTimeout)
	defer cancel()

	res, err := service.userRepository.Create(ctx, userCreate)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (service *authService) IsDuplicateUsername(username string) bool {
	res := service.userRepository.IsDuplicateEmail(username)
	return !(res.Error == nil)
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
