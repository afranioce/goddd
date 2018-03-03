package domain

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
		Check() error
	}
)
