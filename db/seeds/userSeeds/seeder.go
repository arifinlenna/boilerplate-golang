package seeds

import (
	"github.com/go-faker/faker/v4"
	usermodel "github.com/lenna-ai/azureOneSmile.git/db/models/UserModel"
	"gorm.io/gorm"
)

// SeedUsers seeds the database with fake user data
func SeedUsers(db *gorm.DB, count int) error {
	for i := 0; i < count; i++ {
		intRandom,_:=faker.RandomInt(1, 1000)
		user := usermodel.User{
			UserId:    intRandom[0], // generates a random int between 1 and 1000
			Username:  faker.Username(),            // generates a random username
		}
		
		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}
	return nil
}
