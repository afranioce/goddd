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

	assert.Equal(t, dom.Id(), 0)

	err := r.Save(dom)

	id := dom.Id()

	assert.NoError(t, err)
	assert.NotEqual(t, id, 0)

	d, _ := r.First(id)

	assert.Equal(t, d.Id(), id)
}
