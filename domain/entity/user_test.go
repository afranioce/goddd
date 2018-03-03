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

	assert.Equal(t, dom.Id(), uint(0))
	assert.Equal(t, dom.Username(), "username")
	assert.Equal(t, dom.Email(), "email@email.com")
	assert.NotEmpty(t, dom.Password())
	assert.Nil(t, dom.LastLogin())
	assert.Equal(t, dom.Status(), StatusEnabled)

	assert.IsType(t, dom.ToEntity(), &User{})
	assert.IsType(t, dom.ToEntity().ToDomain(), &userDomain{})
}

func BenchmarkUser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dom := NewUser("username", "email@email.com", "123456")
		dom.Id()
		dom.Username()
		dom.Email()
		dom.Password()
		dom.Status()
		dom.LastLogin()
		//Para converter uma entidade para um domínio fica mais rápido fazer o "type assertion" para a entidade. Pq? n sei hsausasau
		e := dom.ToEntity().(*User)
		e.ToDomain()
	}
}
