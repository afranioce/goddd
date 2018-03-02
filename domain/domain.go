package domain

import (
	"gopkg.in/go-playground/validator.v8"
)

type (
	IAggregateRoot interface {
		RootId() uint
	}

	Identitier interface {
		Id() uint
	}

	IValueObject interface {
		Equal(interface{}) bool
	}

	Checker interface {
		Check() validator.ValidationErrors
	}
)
