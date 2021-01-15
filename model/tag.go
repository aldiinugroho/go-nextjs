package model

// GetTagData for ....
type GetTagData []struct {
	Color    string
	Tag_name string
}

// Tag for ....
type Tag struct {
	TagID    uint `gorm:"primaryKey; not null"`
	TagColor string
	TagName  string
}
