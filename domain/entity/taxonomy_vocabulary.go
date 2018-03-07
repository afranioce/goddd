package entity

type TaxonomyVocabulary struct {
	entityBase
	entityBlamed
	Name        string `gorm:"type:varchar(50);not null" sql:"index" check:"required,max=50"`
	Description string `gorm:"type:varchar(1000);not null" check:"max=1000"`
	Status      Status
}

func NewTaxonomyVocabulary(name string, description string, author *User) *TaxonomyVocabulary {
	return &TaxonomyVocabulary{
		Name:        name,
		Description: description,
		Status:      StatusEnabled,
		entityBlamed: entityBlamed{
			CreatedBy:   *author,
			CreatedByID: author.ID,
		},
	}
}

func (d *TaxonomyVocabulary) Check() error {
	return validate.Struct(d)
}
