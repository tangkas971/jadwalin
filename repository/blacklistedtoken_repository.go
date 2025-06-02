package repository

import (
	"jadwalin/model"

	"gorm.io/gorm"
)

type BlacklistTokenRepository interface {
	Save(token model.BlacklistedToken) error
	IsBlacklisted(tokenString string) (bool, error)
}

type blacklistTokenRepository struct{
	db *gorm.DB
}

func NewBlacklistTokenRepository(db *gorm.DB)BlacklistTokenRepository{
	return &blacklistTokenRepository{
		db: db,
	}
}

func (r *blacklistTokenRepository) Save(token model.BlacklistedToken) error {
	return r.db.Save(&token).Error
}

func (r *blacklistTokenRepository) IsBlacklisted(tokenString string)(bool, error){
	var token model.BlacklistedToken
	err := r.db.Where("token = ?", tokenString).First(&token).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound{
			return false, nil 
		}
		return false, err
	}

	return true, nil 
}	