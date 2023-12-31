package service

import (
	"context"
	models "legocy-go/internal/domain/users/models"
	repo "legocy-go/internal/domain/users/repository"
)

type UserImageUseCase struct {
	repo repo.UserImageRepository
}

func NewUserImageUseCase(repo repo.UserImageRepository) UserImageUseCase {
	return UserImageUseCase{repo: repo}
}

func (s *UserImageUseCase) StoreUserImage(c context.Context, image *models.UserImage) error {
	return s.repo.AddUserImage(c, image)
}

func (s *UserImageUseCase) GetUserImages(c context.Context, userID int) ([]*models.UserImage, error) {
	return s.repo.GetUserImages(c, userID)
}
