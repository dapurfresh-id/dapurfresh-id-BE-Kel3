package repositories

import (
	"context"
	"log"

	"github.com/aldisaputra17/dapur-fresh-id/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *entities.User) (*entities.User, error)
	Update(ctx context.Context, user *entities.User) (*entities.User, error)
	FindById(ctx context.Context, id string) ([]*entities.User, error)
	// Image(user *entities.User) (string, error)
	VerifyCredential(username string, password string) interface{}
	IsDuplicateUsername(username string) (tx *gorm.DB)
	FindByUsername(username string) *entities.User
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
	db.connection.Preload("Images").Find(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func (db *userConnection) Update(ctx context.Context, user *entities.User) (*entities.User, error) {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser entities.User
		db.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}

	res := db.connection.WithContext(ctx).Joins("Image").Save(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func (db *userConnection) VerifyCredential(username string, password string) interface{} {
	var user entities.User
	res := db.connection.Where("username = ?", username).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateUsername(username string) (tx *gorm.DB) {
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

func (db *userConnection) FindByUsername(username string) *entities.User {
	var user *entities.User
	db.connection.Where("username = ?", username).Take(&user)
	return user
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
