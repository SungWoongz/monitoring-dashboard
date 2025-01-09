package controller

import (
	"server/models"

	"github.com/shirou/gopsutil/mem"
)

func GetMemStatus() (models.MEM_Status, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return models.MEM_Status{}, err
	}
	vms := &models.MEM_Status{
		VirtualMemoryStat: v,
	}
	return *vms, nil
}
