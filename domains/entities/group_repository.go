package entities

import (
	"context"
	"errors"

	customerrors "github.com/Nexters/pinterest/domains/errors"
	"gorm.io/gorm"
)

type FilmRepository struct {
	*gorm.DB
}

func NewFilmRepository(db *gorm.DB) *FilmRepository {
	return &FilmRepository{db}
}

func (gr *FilmRepository) FindFilm(ctx context.Context, filmId uint) (film Film, err error) {
	err = gr.DB.Preload("PhotoCuts").First(&film, filmId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = customerrors.NewNotFoundError("Film")
			return
		}
		return
	}
	return
}

func (gr *FilmRepository) CountOrderByUserId(ctx context.Context, userId string) (count int64, err error) {
	tx := gr.DB.Model(&Film{}).Where("user_id = ?", userId).Count(&count)
	if tx.Error != nil {
		return
	}
	return
}

func (gr *FilmRepository) SaveFilm(ctx context.Context, film Film) (Film, error) {
	tx := gr.DB.Create(&film)
	if tx.Error != nil {
		return film, customerrors.NewCreateFailedError("Film")
	}
	return film, nil
}

func (gr *FilmRepository) FindAllFilmsInOrder(ctx context.Context, userId string) (films []Film, err error) {
	tx := gr.DB.Where("user_id = ?", userId).Order("order DESC").Find(&films)
	if tx.Error != nil {
		err = tx.Error
		return
	}
	return
}
