package models

type Post struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"-"`
	Slug string `gorm:"unique;notNull" json:"slug"`

	// Content
	Title       string `gorm:"notNull" json:"title"`
	Description string `gorm:"notNull" json:"description"`
	Content     string `gorm:"notNull" json:"content"`
	Tumbnail    string `gorm:"notNull" json:"tumbnail"`

	// Metadata
	CategoryID uint     `gorm:"notNull" json:"-"`
	Category   Category `gorm:"foreignKey:CategoryID" json:"category"`

	Status     string `gorm:"notNull" json:"status"`
	Updated_At int64  `gorm:"autoUpdateTime:milli" json:"update_at"`
	Created_At int64  `gorm:"autoCreateTime" json:"create_at"`
}
