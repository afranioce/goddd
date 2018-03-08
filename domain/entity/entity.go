package entity

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v8"
)

// Status Status of entity
type Status byte

const (
	// StatusExcluded entity excluded
	StatusExcluded Status = iota
	// StatusCanceled entity canceled
	StatusCanceled
	// StatusDisabled entity disabled
	StatusDisabled
	// StatusEnabled entity heabled
	StatusEnabled
)

// Base a base entity definition
type Base struct {
	gorm.Model
}

// IsNew Entity is new
func (d *Base) IsNew() bool {
	return d.ID == 0
}

// Blamed Define the user blame
type Blamed struct {
	CreatedBy   User  `gorm:"save_associations:false"`
	CreatedByID uint  `gorm:"not null"`
	ChangedBy   *User `gorm:"save_associations:false"`
	ChangedByID sql.NullInt64
}

var validate *validator.Validate

func init() {
	config := &validator.Config{TagName: "check"}
	validate = validator.New(config)
}
