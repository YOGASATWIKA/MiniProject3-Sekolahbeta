package model

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_At"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}