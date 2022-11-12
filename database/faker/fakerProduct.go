package faker

// import (
// 	"fmt"
// 	"log"
// 	"math/rand"
// 	"path/filepath"

// 	"github.com/aldisaputra17/dapur-fresh-id/entities"
// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// )

// var nameProd = []string{
// 	"kangkung", "wortel", "bayam", "timun", "tomat", "bumbu rendang", "kacang", "kol", "brokoli",
// 	"kentang", "jipang", "kunyit", "daun pepaya",
// }

// func nameprod() string {
// 	var value string
// 	for _, v := range nameProd {
// 		value = v
// 	}
// 	return value
// }

// var prices = []int{
// 	5000, 1000, 2000, 9000, 10000, 6500, 7000,
// }

// func price() int {
// 	value := 0
// 	for _, v := range prices {
// 		value = prices[v]
// 	}
// 	return value
// }

// func imageProd() string {
// 	file := filepath.Dir("product")
// 	return file
// }

// func ProductFaker(db *gorm.DB) *entities.Product {
// 	id, err := uuid.NewRandom()
// 	if err != nil {
// 		fmt.Println("error uuid")
// 	}

// 	ctg := CategoryFaker(db)
// 	err = db.Create(&ctg).Error
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return &entities.Product{
// 		ID:         id,
// 		Name:       nameprod(),
// 		Price:      price(),
// 		Unit:       rand.Intn(100),
// 		UnitType:   "gr",
// 		Image:      imageProd(),
// 		CategoryID: ctg.ID.String(),
// 	}

// }
