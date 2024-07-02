package models

import (
	"time"
)

type Bike struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Model       string    `json:"model"`
	Manufacturer string   `json:"manufacturer"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
