package controller

import (
	"server/models"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func GetCpuStatus() (models.CPU_Status, error) {
	v := models.CPU_Status{}
	cpuPercent, err := cpu.Percent(1*time.Second, true)
	if err != nil {
		return v, err
	}
	v.Used = cpuPercent
	return v, nil
}
