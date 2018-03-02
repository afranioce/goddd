package entity

import (
	"testing"

	"github.com/afranioce/goddd/domain"
	"github.com/stretchr/testify/assert"
)

func TestTaxonomyVocabularyEmptyInvalid(t *testing.T) {
	user := NewUser("teste", "email@email.com", "123456")
	dom := NewTaxonomyVocabulary("teste", "descricao", user)

	assert.NotNil(t, dom)
}

func TestTaxonomyVocabularyEmpty(t *testing.T) {
	user := NewUser("teste", "email@email.com", "123456")
	dom := NewTaxonomyVocabulary("teste", "descricao", user)

	assert.Equal(t, dom.Id(), uint(0))
	assert.Equal(t, dom.Name(), "teste")
	assert.Equal(t, dom.Description(), "descricao")
	assert.Equal(t, dom.Status(), domain.StatusEnabled)
	assert.IsType(t, dom.ToEntity(), &TaxonomyVocabulary{})
	assert.IsType(t, dom.ToEntity().ToDomain(), &taxonomyVocabularyDomain{})

	assert.IsType(t, dom.Author(), &userDomain{})
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
		//Em questao de performance, fica muito mais rapido converter para dominio seo tipo
		e := dom.ToEntity().(*TaxonomyVocabulary)
		e.ToDomain()
	}
}
