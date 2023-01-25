package models

type Album struct {
	ID     string  `json:"Id"`
	Title  string  `json:"Title"`
	Artist string  `json:"Artist"`
	Price  float32 `json:"Price"`
}
