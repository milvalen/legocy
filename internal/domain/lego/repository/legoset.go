package repository

import (
	"context"
	models "legocy-go/internal/domain/lego/models"
)

type LegoSetRepository interface {
	CreateLegoSet(c context.Context, s *models.LegoSetValueObject) error
	GetLegoSets(c context.Context) ([]*models.LegoSet, error)
	GetLegoSetByID(c context.Context, id int) (*models.LegoSet, error)
	DeleteLegoSet(c context.Context, id int) error
}
