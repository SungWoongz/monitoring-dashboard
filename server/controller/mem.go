package controller

import (
	"log"
	"server/models"
	"time"

	"github.com/shirou/gopsutil/mem"
	"gorm.io/gorm"
)

func GetMemCurrentStatus() (models.MEM_Status, error) {
	v := models.MEM_Status{}
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return v, err
	}
	v.Total = vmStat.Total
	v.Available = vmStat.Available
	v.Used = vmStat.Used
	v.UsedPercent = vmStat.UsedPercent
	v.Timestamp = time.Now()
	return v, nil
}

func GetMemStatus(tx *gorm.DB, interval int, limit int) ([]models.MEM_Status, error) {
	var d []models.MEM_Status
	query := `
		SELECT *
		FROM mem_statuses
		WHERE CAST(strftime('%s', timestamp) AS INTEGER) % ? = 0
		ORDER BY timestamp DESC
		LIMIT ?`
	if err := tx.Raw(query, interval, limit).Scan(&d).Error; err != nil {
		return d, err
	}
	return d, nil
}

func CollectMemStatus(db *gorm.DB) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		memStatus, err := GetMemCurrentStatus()
		if err != nil {
			log.Printf("Failed to get memory status: %v", err)
			continue
		}
		if err := db.Create(&memStatus).Error; err != nil {
			log.Printf("Failed to insert memory status: %v", err)
			continue
		}
		if err := db.Where("timestamp < ?", time.Now().Add(-7*24*time.Hour)).Delete(&models.MEM_Status{}).Error; err != nil {
			log.Printf("Failed to delete old memory records: %v", err)
		}
	}
}
