package entity

import (
	"time"

	"github.com/afranioce/goddd/domain/eventsourcing/event"

	"github.com/afranioce/goddd/domain/eventsourcing"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	eventsourcing.EventSource
	Base
	Username            string `gorm:"type:varchar(150);not null" check:"required"`
	Email               string `gorm:"type:varchar(150);not null;unique" check:"required,email"`
	Password            string `gorm:"type:varchar(255);not null" check:"required"`
	ConfirmationToken   string `gorm:"type:varchar(255);"`
	PasswordRequestedAt time.Time
	LastLogin           *time.Time
	Status              Status
}

func NewUser(username string, email string, plainPassword string) *User {
	usr := &User{}

	eventsourcing.LoadFromEvents(usr, []eventsourcing.Event{
		event.AccountCreated{
			Username: username,
			Email:    email,
		},
		event.PasswordChanged{
			Password: plainPassword,
		},
	})

	return usr
}

func (d *User) UpdateLastLogin() {
	now := time.Now()
	d.LastLogin = &now
}

func (d *User) ComparePassword(plainPassword string) bool {
	byteHash := []byte(d.Password)

	if err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPassword)); err != nil {
		return false
	}

	return true
}

func (d *User) Check() error {
	return validate.Struct(d)
}

func (u *User) Transition(evt eventsourcing.Event) {
	switch e := evt.(type) {
	case event.AccountCreated:
		u.Username = e.Username
		u.Email = e.Email
		u.Status = StatusEnabled
		break
	case event.PasswordChanged:
		if hash, err := bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.MinCost); err == nil {
			u.Password = string(hash)
			u.PasswordRequestedAt = time.Now()
		}
		break
	}
}
