package repository

import (
	"log"
	"os"

	"github.com/afranioce/goddd/config"
	"github.com/afranioce/goddd/domain/entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// InitDB Initialize connection with database
func InitDB() *gorm.DB {
	db, err := gorm.Open(config.MustEnv("DB_DRIVER"), os.ExpandEnv(config.MustEnv("DB_HOST")))
	if err != nil {
		log.Printf("Failed to open connection. host=%s, driver=%s", config.MustEnv("DB_HOST"), config.MustEnv("DB_DRIVER"))
		panic(err)
	}

	// Disable table name's pluralization globally
	db.SingularTable(true)

	if err = db.DB().Ping(); err != nil {
		log.Print(err.Error())
	}

	db.LogMode(true)

	db.AutoMigrate(entity.TaxonomyTerm{}, entity.TaxonomyVocabulary{})

	return db
}
