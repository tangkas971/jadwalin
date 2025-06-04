package repository

import (
	"jadwalin/model"

	"gorm.io/gorm"
)

type ProdiRepository interface {
	CreateProdi(prodi *model.Prodi) error
	FindByCode(code string)(*model.Prodi, error)
	FindyById(id int)(*model.Prodi, error)
	GetAll()([]*model.Prodi, error)
	Delete(id int) error
	Update(prodi *model.Prodi) error
}

type prodiRepository struct {
	db *gorm.DB
}

func NewProdiRepository(db *gorm.DB) ProdiRepository {
	return &prodiRepository{
		db: db,
	}
}

func (r *prodiRepository) CreateProdi(prodi *model.Prodi) error {
	err := r.db.Create(prodi).Error
	return err
}

func (r *prodiRepository) FindByCode(code string)(*model.Prodi, error) {
	var prodi model.Prodi
	err := r.db.Where("code = ?",code).First(&prodi).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			return nil, nil 
		}
		return nil, err
	}

	return &prodi, nil
}

func (r *prodiRepository) FindyById(id int)(*model.Prodi, error){
	var prodi model.Prodi

	err := r.db.Where("id = ?", id).First(&prodi).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			return nil, nil
		}
		return nil, err
	}

	return &prodi, nil 
}

func (r *prodiRepository) GetAll()([]*model.Prodi, error){
	var prodis []*model.Prodi
	err := r.db.Find(&prodis).Error
	if err != nil {
		return nil, err
	}

	return prodis, nil 
}

func (r *prodiRepository) Delete(id int) error{
	err := r.db.Delete(&model.Prodi{}, id).Error
	return err
}

func (r *prodiRepository) Update(prodi *model.Prodi) error{
	err := r.db.Save(prodi).Error
	return err
}

