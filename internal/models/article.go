package models

import "time"

//Article represent article in blog
type Article struct {
	ID          int
	Title       string
	Information string
	Created     time.Time `gorm:"autoCreateTime"`
}
