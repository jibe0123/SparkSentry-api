package entities

import "gorm.io/gorm"

type User struct {
	ID        uint    `gorm:"primaryKey"`
	AccountID *uint   `gorm:"column:account_id"`
	Email     string  `gorm:"size:255;unique;not null"`
	Password  string  `gorm:"size:255;not null"`
	FirstName string  `gorm:"size:100"`
	LastName  string  `gorm:"size:100"`
	Role      string  `gorm:"size:50"` // Ex: Admin, Manager, Viewer
	Account   Account `gorm:"foreignKey:AccountID"`
	gorm.Model
}
