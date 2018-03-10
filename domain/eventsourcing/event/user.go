package event

type AccountCreated struct {
	Username string
	Email    string
	Password string
}

type PasswordChanged struct {
	Password string
}
