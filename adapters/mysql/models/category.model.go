package models

type Category struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"-"`
	Slug string `gorm:"unique;notNull" json:"slug"`

	ParentID uint   `gorm:"default:0" json:"-"`
	Name     string `gorm:"unique;notNull" json:"name"`
}
