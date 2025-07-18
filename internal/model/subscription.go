package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Subscription struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	ServiceName string         `json:"service_name"`
	Price       int            `json:"price"`
	UserID      uuid.UUID      `json:"user_id"`
	StartDate   time.Time      `json:"start_date"`
	EndDate     *time.Time     `json:"end_date,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
