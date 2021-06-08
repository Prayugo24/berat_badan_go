package models

import( 
	"time"
)

type Tb_Berat struct {
	Id int `json:"id", gorm:"primary_key", gorm:autoIncrement` // id
	Tanggal time.Time `json:"tanggal"`
	Max int `json:"max"` 
	Min int `json:"min"` 
	Perbedaan int `json:"perbedaan", gorm:"default:0"` 
}