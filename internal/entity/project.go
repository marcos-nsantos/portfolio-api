package entity

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	ID          uint64 `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:varchar(255);not null"`
	URL         string `gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	UserID      uint64
}
