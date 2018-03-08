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

type Base struct {
	gorm.Model
}

func (d *Base) Id() uint {
	return d.ID
}

func (d *Base) IsNew() bool {
	return d.ID == 0
}

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
