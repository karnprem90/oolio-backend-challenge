package service

import (
	"context"

	"oolio-backend-challenge/internal/domain"
)

type PromoCodeService struct {
	repo domain.PromoRepository
}

func NewPromoCodeService(repo domain.PromoRepository) *PromoCodeService {
	return &PromoCodeService{repo: repo}
}

func (s *PromoCodeService) ValidatePromoCode(ctx context.Context, code string) (bool, error) {
	return s.repo.ValidatePromoCode(ctx, code)
}
