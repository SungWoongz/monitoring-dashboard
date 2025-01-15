package controller

import (
	"log"
	"server/models"
	"time"

	"github.com/shirou/gopsutil/net"
	"gorm.io/gorm"
)

func GetNetCurrentStatus() (models.NET_Status, error) {
	v := models.NET_Status{}
	netIOStats, err := net.IOCounters(false)
	if err != nil {
		return v, err
	}

	if len(netIOStats) > 0 {
		stats := netIOStats[0]
		v.BytesSent = stats.BytesSent
		v.BytesRecv = stats.BytesRecv
		v.PacketsSent = stats.PacketsSent
		v.PacketsRecv = stats.PacketsRecv
		v.Errin = stats.Errin
		v.Errout = stats.Errout
		v.Dropin = stats.Dropin
		v.Dropout = stats.Dropout
		v.Fifoin = stats.Fifoin
		v.Fifoout = stats.Fifoout
		v.Timestamp = time.Now()
	}
	return v, nil
}

func GetNetStatus(tx *gorm.DB, interval int, limit int) ([]models.NET_Status, error) {
	var d []models.NET_Status
	query := `
		SELECT *
		FROM net_statuses
		WHERE CAST(strftime('%s', timestamp) AS INTEGER) % ? = 0
		ORDER BY timestamp DESC
		LIMIT ?`
	if err := tx.Raw(query, interval, limit).Scan(&d).Error; err != nil {
		return d, err
	}
	return d, nil
}

func CollectNetStatus(db *gorm.DB) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		netStatus, err := GetNetCurrentStatus()
		if err != nil {
			log.Printf("Failed to get network status: %v", err)
			continue
		}

		// 데이터베이스에 저장
		if err := db.Create(&netStatus).Error; err != nil {
			log.Printf("Failed to insert network status: %v", err)
			continue
		}

		// 오래된 데이터 삭제 (7일 기준)
		if err := db.Where("timestamp < ?", time.Now().Add(-7*24*time.Hour)).Delete(&models.NET_Status{}).Error; err != nil {
			log.Printf("Failed to delete old network records: %v", err)
		}
	}
}
