package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	entityBase
	Username            string `gorm:"type:varchar(150);not null" check:"required"`
	Email               string `gorm:"type:varchar(150);not null;unique" check:"required,email"`
	Password            string `gorm:"type:varchar(255);not null" check:"required"`
	ConfirmationToken   string `gorm:"type:varchar(255);"`
	PasswordRequestedAt time.Time
	LastLogin           *time.Time
	Status              Status
}

func NewUser(username string, email string, plainPassword string) *User {
	impl := &User{
		Username: username,
		Email:    email,
		Status:   StatusEnabled,
	}

	impl.UpdatePassword(plainPassword)

	return impl
}

func (d *User) UpdateLastLogin() {
	now := time.Now()
	d.LastLogin = &now
}

func (d *User) UpdatePassword(plainPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.MinCost)

	if err == nil {
		d.Password = string(hash)
		d.PasswordRequestedAt = time.Now()
	}

	return err
}

func (d *User) Check() error {
	return validate.Struct(d)
}
