package controller

import (
	"server/models"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func GetCpuStatus() (models.CPU_Status, error) {
	v := models.CPU_Status{}
	cp, err := cpu.Percent(1*time.Second, true)
	if err != nil {
		return v, err
	}
	v.Used = cp
	return v, nil
}
