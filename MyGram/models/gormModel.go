package models

import (
	"time"
)

type GormModel struct {
	ID        uint       `gorm:"PrimaryKey" json:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type PModel struct {
	ID        uint       `gorm:"PrimaryKey" json:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type CModel struct {
	ID        uint       `gorm:"PrimaryKey" json:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type SModel struct {
	ID        uint       `gorm:"PrimaryKey" json:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
