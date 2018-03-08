package entity

// TaxonomyTerm Term of taxonomy
type TaxonomyTerm struct {
	Base
	Blamed
	Name         string             `gorm:"type:varchar(50);not null" check:"required,max=50"`
	Vocabulary   TaxonomyVocabulary `gorm:"save_association:false"`
	VocabularyID uint               `gorm:"not null"`
	Status       Status
}

// NewTaxonomyTerm Constructor
func NewTaxonomyTerm(name string, vocabulary *TaxonomyVocabulary, author *User) *TaxonomyTerm {
	return &TaxonomyTerm{
		Name:         name,
		Status:       StatusEnabled,
		Vocabulary:   *vocabulary,
		VocabularyID: vocabulary.ID,
		Blamed: Blamed{
			CreatedBy:   *author,
			CreatedByID: author.ID,
		},
	}
}

// Check check validation
func (d *TaxonomyTerm) Check() error {
	return validate.Struct(d)
}
