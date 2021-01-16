package models

// GetTagData for ....
type GetTagData []struct {
	Color    string
	Tag_name string
}

// Tag for ....
type Tag struct {
	TagID    uint   `gorm:"primaryKey; not null" json:"TagID"`
	TagColor string `json:"TagColor"`
	TagName  string `json:"TagName"`
}
