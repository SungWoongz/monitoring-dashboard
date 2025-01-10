package models

import "time"

type CPU_Status struct {
	Timestamp time.Time `json:"timestamp" gorm:"primaryKey"`
	Usage     string    `json:"Usage"  gorm:"type:json"`
}
