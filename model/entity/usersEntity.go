package entity

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Address        string `json:"address"`
	Password       string `json:"password"`
	Profilepicture string `json:"profilePicture"`

	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
