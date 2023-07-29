package dto

import (
	"github.com/Nexters/pinterest/domains/entities"
	"github.com/go-playground/validator"
)

func ToFilmDtoList(FilmList []entities.Film) (FilmDtoList []Film, err error) {
	for _, film := range FilmList {
		FilmInfo := Film{
			Title:         film.Title,
			Order:         film.Order,
			PhotoCutCount: film.PhotoCutCount,
			Likes:         film.Likes,
			UserID:        film.UserID,
		}
		validate := validator.New()
		err := validate.Struct(FilmInfo)
		if err != nil {
			return FilmDtoList, err
		}
		FilmDtoList = append(FilmDtoList, FilmInfo)
	}
	return
}

func ToVisitLogDtoList(visitLogList []entities.VisitLog) (visitLogDtoList []VisitLog, err error) {
	for _, visitLog := range visitLogList {
		visitLogInfo := VisitLog{
			UserID: visitLog.UserID,
			Name:   visitLog.Name,
			Text:   visitLog.Text,
		}
		validate := validator.New()
		err := validate.Struct(visitLogInfo)
		if err != nil {
			return visitLogDtoList, err
		}
		visitLogDtoList = append(visitLogDtoList, visitLogInfo)
	}
	return
}

func ToItemDtoList(itemList []entities.Item) (itemDtoList []ItemDetailResponse, err error) {
	for _, item := range itemList {
		itemDetail := ItemDetailResponse{
			Title:     item.Title,
			Text:      item.Text,
			Link:      item.Link,
			Image:     item.Image,
			Likes:     item.Likes,
			GroupID:   item.GroupID,
			CreatedAt: item.CreatedAt,
		}
		itemDtoList = append(itemDtoList, itemDetail)
	}
	return
}
