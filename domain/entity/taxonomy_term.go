package entity

type TaxonomyTerm struct {
	Base
	Blamed
	Name         string             `gorm:"type:varchar(50);not null" check:"required,max=50"`
	Vocabulary   TaxonomyVocabulary `gorm:"save_association:false"`
	VocabularyID uint               `gorm:"not null"`
	Status       Status
}

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

func (d *TaxonomyTerm) Check() error {
	return validate.Struct(d)
}
