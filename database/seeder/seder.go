package seeder

// import (
// 	"github.com/aldisaputra17/dapur-fresh-id/database/faker"
// 	"gorm.io/gorm"
// )

// type Seeder struct {
// 	Seeder interface{}
// }

// func RegisterSeeder(db *gorm.DB) []Seeder {
// 	return []Seeder{
// 		{Seeder: faker.CategoryFaker(db)},
// 		{Seeder: faker.ProductFaker(db)},
// 	}
// }

// func DBSeed(db *gorm.DB) error {
// 	for _, v := range RegisterSeeder(db) {
// 		err := db.Debug().Create(v.Seeder).Error
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
