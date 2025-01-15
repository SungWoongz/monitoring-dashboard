package models

import "time"

type MEM_Status struct {
	Timestamp   time.Time `json:"timestamp" gorm:"primaryKey"` // 기본 키 및 자동 생성 시간
	Total       uint64    `json:"total"`                       // NULL 불가
	Available   uint64    `json:"available"`                   // NULL 불가
	Used        uint64    `json:"used"`                        // NULL 불가
	UsedPercent float64   `json:"usedPercent"`                 // 소수점 2자리 제한
}
