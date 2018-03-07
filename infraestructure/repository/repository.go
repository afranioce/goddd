package repository

import (
	"github.com/afranioce/goddd/domain/entity"
	"github.com/afranioce/goddd/domain"
	"github.com/jinzhu/gorm"
)

type Repository struct {
	*gorm.DB
}

type NewRepository() *Repository {
	return &Repository{
		DB: InitDB()
	}
}

func (r *repositoryImpl) Save(dom domain.Identifier) (err error) {
	if dom.IsNew() == 0 {
		err = r.DB.Create(dom.ToEntity()).Error
	} else {
		err = r.DB.Save(dom.ToEntity()).Error
	}
	return
}