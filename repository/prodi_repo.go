package repository

import (
	"jadwalin/model"

	"gorm.io/gorm"
)

type ProdiRepository interface {
	CreateProdi(prodi *model.Prodi) error
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