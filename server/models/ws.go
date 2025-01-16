package models

type ALL_Status struct {
	CPU CPU_Status `json:"cpu"`
	MEM MEM_Status `json:"mem"`
	NET NET_Status `json:"net"`
}
