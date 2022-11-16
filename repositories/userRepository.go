package repositories

import (
	"context"
	"log"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"github.com/aldisaputra17/dapur-fresh-id/helpers"
	"github.com/aldisaputra17/dapur-fresh-id/request"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *entities.User) (*entities.User, error)
	Update(ctx context.Context, user *entities.User, req *request.RequestUserUpdate) (*entities.User, string, error)
	FindById(ctx context.Context, id string) ([]*entities.User, error)
	// Image(user *entities.User) (string, error)
	VerifyCredential(username string, password string) interface{}
	IsDuplicateEmail(username string) (tx *gorm.DB)
	GetUser(userID string) *entities.User
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) Create(ctx context.Context, user *entities.User) (*entities.User, error) {
	user.Password = hashAndSalt([]byte(user.Password))
	res := db.connection.WithContext(ctx).Create(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func (db *userConnection) Update(ctx context.Context, user *entities.User, req *request.RequestUserUpdate) (*entities.User, string, error) {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser entities.User
		db.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}

	upload, err := helpers.ImageUploadHelper(req.Image)
	if err != nil {
		return nil, "", err
	}

	res := db.connection.WithContext(ctx).Joins("Image").Save(&user)
	if res.Error != nil {
		return nil, "", res.Error
	}
	return user, upload, nil
}

func (db *userConnection) VerifyCredential(username string, password string) interface{} {
	var user entities.User
	res := db.connection.Where("username = ?", username).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateEmail(username string) (tx *gorm.DB) {
	var user entities.User
	return db.connection.Where("username = ?", username).Take(&user)
}

func (db *userConnection) FindById(ctx context.Context, id string) ([]*entities.User, error) {
	var user []*entities.User
	res := db.connection.WithContext(ctx).Where("id = ?", id).Find(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

// func (db *userConnection) Image(user *entities.User) (string, error) {
// 	upload, err := helpers.ImageUploadHelper(user.Image)
// 	if err != nil {
// 		return "", err
// 	}

// 	res := db.connection.Create(&user)
// 	if res.Error != nil {
// 		return "", res.Error
// 	}
// 	return upload, nil
// }

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}

func (db *userConnection) GetUser(userID string) *entities.User {
	log.Println("tes-repo", userID)
	var user *entities.User
	// db.connection.First(&user, userID)
	res := db.connection.Where("id = ?", userID).Take(&user)

	// return user
	if res.Error == nil {
		return user
	}
	return nil
}
