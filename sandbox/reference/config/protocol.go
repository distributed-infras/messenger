package config

type Protocol int

const (
	WebSocket Protocol = iota
	MQTT
	XMPP
	LongPoll
)
