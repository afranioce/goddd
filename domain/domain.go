package domain

type Checker interface {
	Check() error
}

type Identifier interface {
	IsNew() bool
}
