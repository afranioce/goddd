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
	ChangedBy   User `gorm:"save_associations:false"`
	ChangedByID sql.NullInt64
}

type domainBase struct {
	value  DomainTransformer
	errors error
}

func (d *domainBase) Check() validator.ValidationErrors {
	return d.errors.(validator.ValidationErrors)
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
