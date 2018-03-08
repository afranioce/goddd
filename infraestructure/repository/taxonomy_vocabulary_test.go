package repository

import (
	"testing"

	"github.com/afranioce/goddd/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestTaxnomyVocabulary(t *testing.T) {
	author := entity.NewUser("teste", "email@email.com", "123456")
	dom := entity.NewTaxonomyVocabulary("Teste", "", author)
	r := NewRepository()

	assert.Equal(t, dom.ID, uint(0))

	err := r.Save(dom)

	assert.NoError(t, err)
	assert.NotEqual(t, dom.ID, uint(0))

	d := new(entity.TaxonomyVocabulary)
	r.First(d, dom.ID)

	assert.Equal(t, d.ID, dom.ID)

	ds := []entity.TaxonomyVocabulary{}
	r.Find(ds)

	assert.IsType(t, ds, []entity.TaxonomyVocabulary{})
}
