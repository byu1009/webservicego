package models

import "time"

type User struct {
	ID           string     `gorm:"column:id;primaryKey"`
	Password     string     `gorm:"column:password"`
	NIP          string     `gorm:"column:nip"`
	UserAccess   string     `gorm:"column:user_access"`
	GroupAccess  string     `gorm:"column:group_access"`
	PrivateKey   string     `gorm:"column:private_key"`
	PublicKey    string     `gorm:"column:public_key"`
	LastActivity *time.Time `gorm:"column:last_activity"`
	// CreatedAt    time.Time
	// UpdatedAt    time.Time
	// DeletedAt    *time.Time
}

func (User) TableName() string {
	return "io_user"
}
