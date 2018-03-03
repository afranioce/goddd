package entity

type TaxonomyTerm struct {
	entityBase
	entityBlamed
	Name         string             `gorm:"type:varchar(50);not null" check:"required"`
	Vocabulary   TaxonomyVocabulary `gorm:"save_association:false"`
	VocabularyID uint               `gorm:"not null"`
	Status       Status
}

func (e *TaxonomyTerm) ToDomain() EntityTransformer {
	return &taxonomyTermDomain{
		domainBase: &domainBase{
			value: e,
		},
	}
}

type taxonomyTermDomain struct {
	*domainBase
}

func NewTaxonomyTerm(name string, vocabulary *taxonomyVocabularyDomain, author *userDomain) *taxonomyTermDomain {
	return &taxonomyTermDomain{
		domainBase: &domainBase{
			value: &TaxonomyTerm{
				Name:         name,
				Status:       StatusEnabled,
				Vocabulary:   *vocabulary.ToEntity().(*TaxonomyVocabulary),
				VocabularyID: vocabulary.Id(),
				entityBlamed: entityBlamed{
					CreatedBy:   *author.ToEntity().(*User),
					CreatedByID: author.Id(),
				},
			},
		},
	}
}

func (d *taxonomyTermDomain) Id() uint {
	return d.value.(*TaxonomyTerm).ID
}

func (d *taxonomyTermDomain) Name() string {
	return d.value.(*TaxonomyTerm).Name
}

func (d *taxonomyTermDomain) Vocabulary() EntityTransformer {
	return d.value.(*TaxonomyTerm).Vocabulary.ToDomain()
}

func (d *taxonomyTermDomain) Author() EntityTransformer {
	return d.value.(*TaxonomyTerm).CreatedBy.ToDomain()
}

func (d *taxonomyTermDomain) Editor() EntityTransformer {
	return d.value.(*TaxonomyTerm).ChangedBy.ToDomain()
}

func (d *taxonomyTermDomain) Status() Status {
	return d.value.(*TaxonomyTerm).Status
}
