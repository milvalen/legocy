package service

import (
	"golang.org/x/net/context"
	models "legocy-go/internal/domain/lego/models"
	r "legocy-go/internal/domain/lego/repository"
)

type LegoSeriesService struct {
	repo r.LegoSeriesRepository
}

func NewLegoSeriesService(repo r.LegoSeriesRepository) LegoSeriesService {
	return LegoSeriesService{repo: repo}
}

func (s *LegoSeriesService) ListSeries(ctx context.Context) ([]*models.LegoSeries, error) {
	return s.repo.GetLegoSeriesList(ctx)
}

func (s *LegoSeriesService) DetailSeries(ctx context.Context, id int) (*models.LegoSeries, error) {
	return s.repo.GetLegoSeries(ctx, id)
}

func (s *LegoSeriesService) CreateLegoSeries(ctx context.Context, m *models.LegoSeriesValueObject) error {
	return s.repo.CreateLegoSeries(ctx, m)
}

func (s *LegoSeriesService) DeleteSeries(ctx context.Context, id int) error {
	return s.repo.DeleteLegoSeries(ctx, id)
}
