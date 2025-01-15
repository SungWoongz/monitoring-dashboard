package controller

import (
	"server/models"

	"github.com/shirou/gopsutil/host"
)

func GetServerInfo() (models.ServerInfo, error) {
	i, err := host.Info()
	if err != nil {
		return models.ServerInfo{}, err
	}
	rv := models.ServerInfo{
		Hostname:             i.Hostname,
		Uptime:               i.Uptime,
		BootTime:             i.BootTime,
		OS:                   i.OS,
		Platform:             i.Platform,
		PlatformFamily:       i.PlatformFamily,
		PlatformVersion:      i.PlatformVersion,
		KernelVersion:        i.KernelVersion,
		KernelArch:           i.KernelArch,
		VirtualizationSystem: i.VirtualizationSystem,
		VirtualizationRole:   i.VirtualizationRole,
		HostID:               i.HostID,
	}
	return rv, nil
}
