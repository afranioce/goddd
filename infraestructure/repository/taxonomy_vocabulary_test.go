package repository

import (
	"testing"

	"github.com/afranioce/goddd/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestTaxnomyVocabulary(t *testing.T) {
	author := entity.NewUser("teste", "email@email.com", "123456")
	dom := entity.NewTaxonomyVocabulary("Teste", "", author)
	r := NewTaxnomyVocabulary()

	assert.Equal(t, dom.ID, 0)

	err := r.Save(dom)

	assert.NoError(t, err)
	assert.NotEqual(t, dom.ID, 0)

	d, _ := r.First(dom.ID)

	assert.Equal(t, d.ID, dom.ID)
}
