package osinfo

type Memory struct {
	Type  string  `json:"type,omitempty"`
	Speed uint    `json:"speed,omitempty"` // RAM data rate in MT/s
	Total float64 `json:"total,omitempty"` // Total RAM size in GiB
	Used  float64 `json:"used,omitempty"`  // Used RAM size in GiB
}
