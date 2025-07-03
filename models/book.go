package models

type Book struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Quantity  int    `json:"quantity"`
}
