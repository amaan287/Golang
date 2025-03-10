package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct{
	gorm.Model
	Title string
	Description string
	CreatedAt time.Time
}