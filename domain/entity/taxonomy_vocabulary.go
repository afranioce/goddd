package entity

// TaxonomyVocabulary Vocabulary of taxonomy
type TaxonomyVocabulary struct {
	Base
	Blamed
	Name        string `gorm:"type:varchar(50);not null" sql:"index" check:"required,max=50"`
	Description string `gorm:"type:varchar(1000);not null" check:"max=1000"`
	Status      Status
}

// NewTaxonomyVocabulary Constructor
func NewTaxonomyVocabulary(name string, description string, author *User) *TaxonomyVocabulary {
	return &TaxonomyVocabulary{
		Name:        name,
		Description: description,
		Status:      StatusEnabled,
		Blamed: Blamed{
			CreatedBy:   *author,
			CreatedByID: author.ID,
		},
	}
}

// Check check validation
func (d *TaxonomyVocabulary) Check() error {
	return validate.Struct(d)
}
