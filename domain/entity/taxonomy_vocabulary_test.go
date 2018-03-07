package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaxonomyVocabularyCheckSuccess(t *testing.T) {
	user := NewUser("teste", "email@email.com", "123456")
	dom := NewTaxonomyVocabulary("teste", "descricao", user)

	assert.Equal(t, dom.ID, uint(0))
	assert.Equal(t, dom.Name, "teste")
	assert.Equal(t, dom.Description, "descricao")
	assert.Equal(t, dom.Status, StatusEnabled)

	assert.IsType(t, dom.CreatedBy, User{})
	assert.IsType(t, dom.ChangedBy, &User{})

	assert.Nil(t, dom.Check())
}

func TestTaxonomyVocabularyCheckFailure(t *testing.T) {
	user := NewUser("teste", "email@email.com", "123456")

	assert.NotNil(t, NewTaxonomyVocabulary("", "descricao", user).Check())
}

func BenchmarkTaxonomyVocabulary(b *testing.B) {
	user := NewUser("teste", "email@email.com", "123456")
	for i := 0; i < b.N; i++ {
		dom := NewTaxonomyVocabulary("teste", "descricao", user)
		dom.Check()
	}
}
