package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserDomainCheckFailure(t *testing.T) {
	assert.NotNil(t, NewUser("", "email@email.com", "123456").Check())
	assert.NotNil(t, NewUser("teste", "", "123456").Check())
	assert.NotNil(t, NewUser("teste", "emailwer", "").Check())
}

func TestUserDomainCheckSuccess(t *testing.T) {
	dom := NewUser("username", "email@email.com", "123456")

	assert.Equal(t, dom.ID, uint(0))
	assert.Equal(t, dom.Username, "username")
	assert.Equal(t, dom.Email, "email@email.com")
	assert.NotEmpty(t, dom.Password)
	assert.Nil(t, dom.LastLogin)
	assert.Equal(t, dom.Status, StatusEnabled)
}

func BenchmarkUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dom := NewUser("username", "email@email.com", "123456")
		dom.Check()
	}
}
