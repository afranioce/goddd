package entity

type TaxonomyVocabulary struct {
	entityBase
	entityBlamed
	Name        string `gorm:"type:varchar(50);not null" sql:"index" check:"required,max=50"`
	Description string `gorm:"type:varchar(1000);not null" check:"max=1000"`
	Status      Status
}

func (e *TaxonomyVocabulary) ToDomain() EntityTransformer {
	return &TaxonomyVocabularyDomain{
		domainBase: &domainBase{
			value: e,
		},
	}
}

type TaxonomyVocabularyDomain struct {
	*domainBase
}

func NewTaxonomyVocabulary(name string, description string, author *userDomain) *TaxonomyVocabularyDomain {
	return &TaxonomyVocabularyDomain{
		domainBase: &domainBase{
			value: &TaxonomyVocabulary{
				Name:        name,
				Description: description,
				Status:      StatusEnabled,
				entityBlamed: entityBlamed{
					CreatedBy:   *author.ToEntity().(*User),
					CreatedByID: author.Id(),
				},
			},
		},
	}
}

func (d *TaxonomyVocabularyDomain) Id() uint {
	return d.value.(*TaxonomyVocabulary).ID
}

func (d *TaxonomyVocabularyDomain) Name() string {
	return d.value.(*TaxonomyVocabulary).Name
}

func (d *TaxonomyVocabularyDomain) Description() string {
	return d.value.(*TaxonomyVocabulary).Description
}

func (d *TaxonomyVocabularyDomain) Author() EntityTransformer {
	return d.value.(*TaxonomyVocabulary).CreatedBy.ToDomain()
}

func (d *TaxonomyVocabularyDomain) Editor() EntityTransformer {
	return d.value.(*TaxonomyVocabulary).ChangedBy.ToDomain()
}

func (d *TaxonomyVocabularyDomain) Status() Status {
	return d.value.(*TaxonomyVocabulary).Status
}

func (d *TaxonomyVocabularyDomain) ToEntity() DomainTransformer {
	return &d.value
}

