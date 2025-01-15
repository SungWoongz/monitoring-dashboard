package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"server/models"
	"strconv"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"gorm.io/gorm"
)

func GetCpuCurrentStatus() (models.CPU_Status, error) {
	v := models.CPU_Status{}
	cp, err := cpu.Percent(1*time.Second, true)
	if err != nil {
		return v, err
	}
	cpstring, err := json.Marshal(cp)
	if err != nil {
		return v, err
	}
	v.Usage = string(cpstring)
	return v, nil
}

func GetCpuStatus(tx *gorm.DB, interval int, limit int) ([]models.CPU_Status, error) {
	var d []models.CPU_Status
	query := `
		SELECT *
		FROM cpu_statuses
		WHERE CAST(strftime('%s', timestamp) AS INTEGER) % ? = 0
		ORDER BY timestamp DESC
		LIMIT ?`
	if err := tx.Raw(query, interval, limit).Scan(&d).Error; err != nil {
		return d, err
	}
	return d, nil
}

func CollectCpuStatus(db *gorm.DB) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		var err error
		cpuStatus, err := GetCpuCurrentStatus()
		if err != nil {
			log.Printf("Failed to get records: %v", err)
		}
		cpuStatus.Timestamp = time.Now()
		if err := db.Create(&cpuStatus).Error; err != nil {
			log.Printf("Failed to insert CPU usage: %v", err)
		}
		if err := db.Where("timestamp < ?", time.Now().Add(-7*24*time.Hour)).Delete(&models.CPU_Status{}).Error; err != nil {
			log.Printf("Failed to delete old records: %v", err)
		}
	}
}

// [NOTUSE!] ConvertStringToFloatArray converts a string like "[1.23,4.56]" to a float64 array
func ConvertStringToFloatArray(input string) ([]float64, error) {
	input = strings.Trim(input, "[]")
	strValues := strings.Split(input, ",")
	var floatValues []float64
	for _, str := range strValues {
		val, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
		if err != nil {
			return nil, fmt.Errorf("failed to parse '%s' as float: %v", str, err)
		}
		floatValues = append(floatValues, val)
	}
	return floatValues, nil
}
