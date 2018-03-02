package domain

import (
	"gopkg.in/go-playground/validator.v8"
)

type (
	Status byte

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

const (
	StatusExcluded Status = iota
	StatusCanceled
	StatusDisabled
	StatusEnabled
)
