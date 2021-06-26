package models

import "time"

type Image struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
