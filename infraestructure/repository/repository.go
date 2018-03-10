package repository

import (
	"github.com/afranioce/goddd/domain"
	"github.com/jinzhu/gorm"
)

type Repository struct {
	*gorm.DB
}

func NewRepository() *Repository {
	return &Repository{
		DB: InitDB(),
	}
}

func (r *Repository) Save(dom domain.Identifier) error {
	if dom.IsNew() {
		return r.DB.Create(dom).Error
	} else {
		return r.DB.Save(dom).Error
	}
}
