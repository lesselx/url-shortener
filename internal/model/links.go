package data

import (
	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	OriginUrl *string `json:"origin_url"`
	Alias     *string `json:"alias"`
}
