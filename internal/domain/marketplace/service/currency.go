package service

import (
	"context"
	models "legocy-go/internal/domain/marketplace/models"
	"legocy-go/internal/domain/marketplace/repository"
)

type CurrencyUseCase struct {
	repo marketplace.CurrencyRepository
}

func NewCurrencyUseCase(repo marketplace.CurrencyRepository) CurrencyUseCase {
	return CurrencyUseCase{repo: repo}
}

func (s CurrencyUseCase) CurrenciesList(c context.Context) ([]*models.Currency, error) {
	return s.repo.GetCurrencies(c)
}

func (s CurrencyUseCase) CurrencyDetail(c context.Context, symbol string) (*models.Currency, error) {
	return s.repo.GetCurrency(c, symbol)
}

func (s CurrencyUseCase) CreateCurrency(c context.Context, curr *models.CurrencyValueObject) error {
	return s.repo.CreateCurrency(c, curr)
}
