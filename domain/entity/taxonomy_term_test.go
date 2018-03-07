package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaxonomyTermValid(t *testing.T) {
	user := NewUser("teste", "email@email.com", "123456")
	vocabDom := NewTaxonomyVocabulary("teste", "descricao", user)
	dom := NewTaxonomyTerm("termname", vocabDom, user)

	assert.Equal(t, dom.ID, uint(0))
	assert.Equal(t, dom.Name, "termname")
	assert.Equal(t, dom.Status, StatusEnabled)

	assert.IsType(t, dom.Vocabulary, TaxonomyVocabulary{})

	assert.Nil(t, dom.Check())
}

func TestTaxonomyTermCheckFailure(t *testing.T) {
	user := NewUser("teste", "email@email.com", "123456")
	vocabDom := NewTaxonomyVocabulary("teste", "descricao", user)

	assert.NotNil(t, NewTaxonomyTerm("", vocabDom, user).Check())
}

func BenchmarkTaxonomyTerm(b *testing.B) {
	user := NewUser("teste", "email@email.com", "123456")
	vocabDom := NewTaxonomyVocabulary("teste", "descricao", user)
	for i := 0; i < b.N; i++ {
		dom := NewTaxonomyTerm("termname", vocabDom, user)
		dom.Check()
	}
}
