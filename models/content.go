package models

import "time"

// GetContentData for ....
type GetContentData []struct {
	Content string
	Tag     string
	Title   string
}

// Content for ....
type Content struct {
	ContentID      uint      `gorm:"primaryKey; not null" json:"ContentID"`
	TagID          int       `gorm:"not null" json:"TagID"`
	Tag            Tag       `gorm:"ForeignKey:TagID;References:TagID" json:"Tag"`
	ContentContent string    `json:"ContentContent"`
	ContentTitle   string    `json:"ContentTitle"`
	CreatedAt      time.Time `json:"CreatedAt"`
	UpdatedAt      time.Time `json:"UpdatedAt"`
}
