package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int            `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" form:"name" validate:"gte=6,lte=32" gorm:"not null"`
	Email     string         `json:"email" form:"email" validate:"required,email" gorm:"not null"`
	Password  string         `json:"-" form:"password" validate:"required,gte=8" gorm:"not null,colum:password"`
	Phone     int            `json:"phone" form:"phone" validate:"required,number,min=12" gorm:"required,not null"`
	Role    string `json:"role" gorm:"default:'user'"`
    JobID   uint   `json:"job_id"`          // Foreign key ke Job
    Job     Job    `json:"job" gorm:"foreignKey:JobID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}