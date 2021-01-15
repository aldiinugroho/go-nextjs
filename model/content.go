package model

// GetContentData for ....
type GetContentData []struct {
	Content string
	Tag     string
	Title   string
}

// Content for ....
type Content struct {
	ContentID      uint `gorm:"primaryKey; not null"`
	TagID          int  `gorm:"not null"`
	Tag            Tag  `gorm:"ForeignKey:TagID;References:TagID"`
	ContentContent string
	ContentTitle   string
}
