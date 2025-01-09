package controller

import (
	"server/models"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func GetCpuStatus() models.CPU_Status {
	v := models.CPU_Status{}
	cpuPercent, _ := cpu.Percent(1*time.Second, false)
	v.Used = cpuPercent
	return v
}
