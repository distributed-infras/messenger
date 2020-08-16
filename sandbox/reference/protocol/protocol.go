package protocol

type Protocol int

const (
	WebSocket Protocol = iota
	MQTT
	XMPP
	LongPoll
)
