package repository

import (
	"context"
	models "legocy-go/internal/domain/lego/models"
)

type LegoSeriesRepository interface {
	CreateLegoSeries(c context.Context, s *models.LegoSeriesValueObject) error
	GetLegoSeriesList(c context.Context) ([]*models.LegoSeries, error)
	GetLegoSeries(c context.Context, id int) (*models.LegoSeries, error)
	GetLegoSeriesByName(c context.Context, name string) (*models.LegoSeries, error)
	DeleteLegoSeries(c context.Context, id int) error
}
