package models

type AuthToken struct {
	Token string `gorm:"primaryKey;notNull" json:"token"`
}
