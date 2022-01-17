package models

import (
	"gorm.io/gorm"
	"time"
)

type ID struct {
	ID uint `json:"id" gorm:"primaryKey"`
}

type Timestamps struct {
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type SoftDeletes struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
