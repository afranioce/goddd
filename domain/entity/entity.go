package entity

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v8"
)

type Status byte

const (
	StatusExcluded Status = iota
	StatusCanceled
	StatusDisabled
	StatusEnabled
)

type entityBase struct {
	gorm.Model
}

func (d *entityBase) Id() uint {
	return d.ID
}

func (d *entityBase) IsNew() bool {
	return d.ID == 0
}

type entityBlamed struct {
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
