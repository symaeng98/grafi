package dto

type PhotoCutUpdateRequest struct {
	ID     uint   `json:"photo_cut_id" validate:"required"`
	Title  string `json:"title"`
	Text   string `json:"text"`
	Image  string `json:"image"`
	FilmID uint   `json:"film_id" validate:"required"`
}
