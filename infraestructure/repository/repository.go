package repository

import (
	"github.com/afranioce/goddd/domain/entity"
	"github.com/jinzhu/gorm"
)

type repositoryImpl struct {
	DB     *gorm.DB
	values interface{}
}

func (r *repositoryImpl) Fetch(limit int, offset int, where ...interface{}) ([]entity.EntityTransformer, error) {
	tmp := r.values.([]entity.DomainTransformer)

	if err := r.DB.Limit(limit).Offset(offset).Find(&tmp, where...).Error; err != nil {
		return nil, err
	}

	doms := make([]entity.EntityTransformer, len(tmp))
	for i, e := range tmp {
		doms[i] = e.ToDomain()
	}
	return doms, nil
}

func (r *repositoryImpl) Save(dom entity.EntityTransformer) (err error) {
	if dom.Id() == 0 {
		err = r.DB.Create(dom.ToEntity()).Error
	} else {
		err = r.DB.Save(dom.ToEntity()).Error
	}
	return
}

func (r *repositoryImpl) First(where ...interface{}) (entity.EntityTransformer, error) {
	tmp := r.values.([]entity.DomainTransformer)[0]
	err := r.DB.First(&tmp, where...).Error
	return tmp.ToDomain(), err
}

func (r *repositoryImpl) Delete(where ...interface{}) error {
	d, err := r.First(where...)
	if err == nil {
		r.DB.Delete(d.ToEntity())
	}
	return err
}
