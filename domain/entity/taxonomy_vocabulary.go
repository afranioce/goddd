package entity

import (
	"github.com/afranioce/goddd/domain"
)

type TaxonomyVocabulary struct {
	entityBase
	Name        string `gorm:"type:varchar(50);not null" sql:"index"`
	Description string `gorm:"type:varchar(1000);not null"`
	CreatedBy   User   `gorm:"save_associations:false"`
	CreatedByID uint   `gorm:"not null"`
	ChangedBy   User   `gorm:"save_associations:false"`
	ChangedByID uint
	Status      domain.Status
}

func (entidade *TaxonomyVocabulary) ToDomain() EntityTransformer {
	return &taxonomyVocabularyDomain{
		domainBase: &domainBase{
			value: entidade,
		},
	}
}

type taxonomyVocabularyDomain struct {
	*domainBase
}

func NewTaxonomyVocabulary(nome string, descricao string, author *userDomain) *taxonomyVocabularyDomain {
	return &taxonomyVocabularyDomain{
		domainBase: &domainBase{
			value: &TaxonomyVocabulary{
				Name:        nome,
				Description: descricao,
				Status:      domain.StatusEnabled,
				CreatedBy:   *author.ToEntity().(*User),
				CreatedByID: author.Id(),
			},
		},
	}
}

func (d *taxonomyVocabularyDomain) Id() uint {
	return d.value.(*TaxonomyVocabulary).ID
}

func (d *taxonomyVocabularyDomain) Name() string {
	return d.value.(*TaxonomyVocabulary).Name
}

func (d *taxonomyVocabularyDomain) Description() string {
	return d.value.(*TaxonomyVocabulary).Description
}

func (d *taxonomyVocabularyDomain) Author() EntityTransformer {
	return d.value.(*TaxonomyVocabulary).CreatedBy.ToDomain()
}

func (d *taxonomyVocabularyDomain) Status() domain.Status {
	return d.value.(*TaxonomyVocabulary).Status
}
