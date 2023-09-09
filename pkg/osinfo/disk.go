package osinfo

type Disk struct {
	Devices string  `json:"devices,omitempty"` // physical disk devices as string
	Total   float64 `json:"total,omitempty"`   // Total Disk size in GiB
	Used    float64 `json:"used,omitempty"`    // Used Disk size in GiB
	Free    float64 `json:"free,omitempty"`    // Free Disk size in GiB
}
