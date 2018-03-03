package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaxonomyVocabularyCheckSuccess(t *testing.T) {
	user := NewUser("teste", "email@email.com", "123456")
	dom := NewTaxonomyVocabulary("teste", "descricao", user)

	assert.Equal(t, dom.Id(), uint(0))
	assert.Equal(t, dom.Name(), "teste")
	assert.Equal(t, dom.Description(), "descricao")
	assert.Equal(t, dom.Status(), StatusEnabled)
	assert.IsType(t, dom.ToEntity(), &TaxonomyVocabulary{})
	assert.IsType(t, dom.ToEntity().ToDomain(), &taxonomyVocabularyDomain{})

	assert.IsType(t, dom.Author(), &userDomain{})
	assert.IsType(t, dom.Editor(), &userDomain{})

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
		dom.Id()
		dom.Name()
		dom.Description()
		dom.Status()
		dom.Author()
		dom.Editor()
		//Para converter uma entidade para um domínio fica mais rápido fazer o "type assertion" para a entidade. Pq? n sei hsausasau
		e := dom.ToEntity().(*TaxonomyVocabulary)
		e.ToDomain()
		dom.Check()
	}
}
