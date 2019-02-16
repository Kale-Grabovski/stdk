package domain

// Request
type InstalledModule struct {
	Id      string `json:"id"`
	Version int    `json:"version"`
}

type InstalledModules struct {
	DeviceId string            `json:"deviceId" binding:"required"`
	Modules  []InstalledModule `json:"installedModules" binding:"required"`
}

// Response
type RevalidateModule struct {
	Id       string `json:"id"`
	Version  int    `json:"version"`
	Settings string `json:"settings"`
}

type RevalidateModules struct {
	DeviceId string             `json:"deviceId"`
	Modules  []RevalidateModule `json:"revalidateModules"`
}
