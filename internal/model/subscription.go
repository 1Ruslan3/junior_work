package model

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	ServiceName string     `json:"service_name"`
	Price       int        `json:"price"`
	UserID      uuid.UUID  `json:"user_id" gorm:"type:uuid"`
	StartDate   time.Time  `json:"start_date"`
	EndDate     *time.Time `json:"end_date,omitempty"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
