package models

import "github.com/shirou/gopsutil/mem"

type MEM_Status struct {
	*mem.VirtualMemoryStat
}
