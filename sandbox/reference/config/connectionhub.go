package config

type ConnectionHub struct {
	Host     string   `json:"ip"`
	Port     int      `json:"port"`
	Protocol Protocol `json:"protocol"`
}
