package models

type User struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"-"`

	// User data
	FirstName string `gorm:"notNull" json:"first_name"`
	LastName  string `gorm:"notNull" json:"last_name"`

	Username string `gorm:"unique;notNull" json:"username"`
	Email    string `gorm:"unique;notNull" json:"email"`
	Password string `gorm:"notNull" json:"-"`
	Role     string `gorm:"notNull" json:"role"` // admin, user
	Avatar   string `gorm:"notNull" json:"avatar"`

	// Metadata
	Status     string `gorm:"notNull" json:"status"`
	Updated_At int64  `gorm:"autoUpdateTime:milli" json:"update_at"`
	Created_At int64  `gorm:"autoCreateTime" json:"create_at"`
}
