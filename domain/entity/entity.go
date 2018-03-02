package entity

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v8"
)

type entityBase struct {
	gorm.Model
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
