package model

type Car struct {
	Id    int    `json:"id"`
	Year  int    `json:"year"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Color string `json:"color"`
}
