package config

import (
	"github.com/Mangaba-Labs/tempoo-api/pkg/domain/user/model"
	"gorm.io/gorm"
	"log"
)

type Migrate struct {
	DB *gorm.DB
}

func (m *Migrate) MigrateAll() (err error) {
	log.Println("Migrating database... 🤞")
	err = m.DB.AutoMigrate(&model.User{})

	if err != nil {
		log.Fatal("Something went wrong on db migration process...\n ", err)
	}

	log.Println("Database migrated with success 😁")
	return
}
