package db

import (
	"log"

	"github.com/prasanth-pn/Movies-onito-test-task/pkg/config"
	"github.com/prasanth-pn/Movies-onito-test-task/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectGorm(conf config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(conf.DBSource), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatalf("error while connecting gorm %s", err)
	}
	db.AutoMigrate(&domain.Ratings{})
	db.AutoMigrate(&domain.Movies{})
	return db
}
