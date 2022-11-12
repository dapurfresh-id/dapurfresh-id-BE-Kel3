package faker

// import (
// 	"fmt"

// 	"github.com/aldisaputra17/dapur-fresh-id/entities"
// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// )

// var nameCategory = []string{
// 	"bumbu", "sayur", "minyak", "sembako", "daging",
// }

// func namecategory() string {
// 	var value string
// 	for _, v := range nameCategory {
// 		value = v
// 	}
// 	return value
// }

// func imageCate() string {
// }

// func CategoryFaker(db *gorm.DB) *entities.Category {
// 	id, err := uuid.NewRandom()
// 	if err != nil {
// 		fmt.Println("error uuid")
// 	}

// 	return &entities.Category{
// 		ID:    id,
// 		Name:  namecategory(),
// 		Image: imageCate(),
// 	}
// }
