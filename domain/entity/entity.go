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

type entityBlamed struct {
	CreatedBy   User `gorm:"save_associations:false"`
	CreatedByID uint `gorm:"not null"`
	ChangedBy   User `gorm:"save_associations:false" check:"required,structonly"`
	ChangedByID sql.NullInt64
}

var validate *validator.Validate

func init() {
	config := &validator.Config{TagName: "check"}
	validate = validator.New(config)
}

type domainBase struct {
	value DomainTransformer
}

func (d *domainBase) Check() error {
	return validate.Struct(d.value)
}

func (d *domainBase) ToEntity() DomainTransformer {
	return d.value
}

type DomainTransformer interface {
	ToDomain() EntityTransformer
}

type EntityTransformer interface {
	ToEntity() DomainTransformer
}
