package models

import "time"

//Article represent article in blog
type Article struct {
	ID          int `gorm:"primaryKey"`
	Title       string
	Information string
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
