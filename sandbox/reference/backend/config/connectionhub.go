package config

import "github.com/distributed-infras/messenger/sandbox/reference/protocol"

type ConnectionHub struct {
	Host     string            `json:"host"`
	Port     int               `json:"port"`
	Protocol protocol.Protocol `json:"protocol"`
}
