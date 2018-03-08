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

func (r *Repository) Save(dom domain.Identifier) (err error) {
	if dom.IsNew() {
		err = r.DB.Create(dom).Error
	} else {
		err = r.DB.Save(dom).Error
	}
	return
}
