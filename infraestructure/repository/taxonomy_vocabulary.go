package repository

import (
	"github.com/afranioce/goddd/domain/entity"
)

func NewTaxnomyVocabulary() *taxonomyVocabularyImpl {
	return &taxonomyVocabularyImpl{
		repositoryImpl: &repositoryImpl{
			DB:     InitDB(),
			values: make([]entity.TaxonomyVocabulary, 0),
		},
	}
}

type taxonomyVocabularyImpl struct {
	*repositoryImpl
}
