package entity

import (
	"time"

	"github.com/afranioce/goddd/domain"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	entityBase
	Username            string `gorm:"type:varchar(150);not null"`
	Email               string `gorm:"type:varchar(150);not null;unique"`
	Password            string `gorm:"type:varchar(255);not null"`
	ConfirmationToken   string `gorm:"type:varchar(255);"`
	PasswordRequestedAt time.Time
	LastLogin           *time.Time
	Status              domain.Status
}

func (entidade *User) ToDomain() EntityTransformer {
	return &userDomain{
		&domainBase{
			value: entidade,
		},
	}
}

func NewUser(username string, email string, plainPassword string) *userDomain {
	impl := &userDomain{
		&domainBase{
			value: &User{
				Username: username,
				Email:    email,
				Status:   domain.StatusEnabled,
			},
		},
	}

	impl.UpdatePassword(plainPassword)

	return impl
}

type userDomain struct {
	*domainBase
}

func (d *userDomain) Id() uint {
	return d.value.(*User).ID
}

func (d *userDomain) Username() string {
	return d.value.(*User).Username
}

func (d *userDomain) Email() string {
	return d.value.(*User).Email
}

func (d *userDomain) Password() string {
	return d.value.(*User).Password
}

func (d *userDomain) LastLogin() time.Time {
	return *d.value.(*User).LastLogin
}

func (d *userDomain) UpdateLastLogin() {
	now := time.Now()
	d.value.(*User).LastLogin = &now
}

func (d *userDomain) UpdatePassword(plainPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.MinCost)

	if err == nil {
		d.value.(*User).Password = string(hash)
		d.value.(*User).PasswordRequestedAt = time.Now()
	}

	return err
}

func (d *userDomain) Status() domain.Status {
	return d.value.(*User).Status
}
