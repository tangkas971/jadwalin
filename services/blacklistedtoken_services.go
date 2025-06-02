package services

import (
	"jadwalin/model"
	"jadwalin/repository"
	"time"
)

type BlacklistTokenService interface {
	BlacklistToken(token string, expiresAt time.Time) error
	IsBlacklisted(token string)(bool, error)
}

type blacklistTokenService struct {
	repo repository.BlacklistTokenRepository
}

func NewBlacklistTokenService(repo repository.BlacklistTokenRepository) BlacklistTokenService {
	return &blacklistTokenService{
		repo: repo,
	}
}

func (s *blacklistTokenService) BlacklistToken(token string, expiresAt time.Time) error {
	blacklisted := model.BlacklistedToken{
		Token: token,
		ExpiresAt: expiresAt,
	}

	return s.repo.Save(blacklisted)
}

func (s *blacklistTokenService) IsBlacklisted(token string)(bool, error){
	return s.repo.IsBlacklisted(token)
}