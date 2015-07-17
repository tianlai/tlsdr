package main
import (
	"container/list"
	"fmt"
)

type Connection struct {
	Success         bool
	Events          list.List
	SrcHost         string
	DestHost        string
	Recommendations list.List
	FailedReason    string

	EventsArray    []Event
	RecommendationsArray []string
}

type Event struct {

	Success   bool
	EventType string
	C2s       bool
	Code      int
}

func NewConnection(from string, to string) Connection {
	return Connection{Success:true, Events:list.List{}, SrcHost:from, DestHost:to}
}

func NewEvent(ucode uint8) Event {
	code := int(ucode)
	// TODO Fix these things
	lookup := make(map[int]string)

	lookup[20] = "TypeChangeCypherSpec"
	lookup[21] = "TypeAlert"
	lookup[22] = "TypeHandshake"
	lookup[23] = "TypeApplicationData"
	lookup[0] = "HandshakeTypeHelloRequest"
	lookup[1] = "HandshakeTypeClientHello"
	lookup[2] = "HandshakeTypeServerHello"
	lookup[3] = "HandshakeTypeHelloVerifyRequest"
	lookup[11] = "HandshakeTypeCertificate"
	lookup[12] = "HandshakeTypeServerKeyExchange"
	lookup[13] = "HandshakeTypeCertificateRequest"
	lookup[14] = "HandshakeTypeServerHelloDone"
	lookup[15] = "HandshakeTypeCertificateVerify"
	lookup[16] = "HandshakeTypeClientKeyExchange"
	lookup[20] = "HandshakeTypeFinished"

	// TODO FIx this
	if code == 22 || code == 11 || code == 13 {
		return Event{Success:true, EventType:lookup[code], C2s: true, Code: code}
	} else {
		return Event{Success:true, EventType:lookup[code], C2s: false, Code: code}
	}
}

func (event Event) String() string {
	return fmt.Sprintf("Event{success: %t, type: '%s'($d)}", event.Success, event.EventType, event.Code)
}

func (connection Connection) String() string {
	return fmt.Sprintf("Connection{success: %t, failReason: '%s', src: '%s', dest: '%s', recommendations: %#v, events: %#v}",
		connection.Success, connection.FailedReason, connection.SrcHost, connection.DestHost, connection.Recommendations, connection.Events)
}

func (connection Connection) AddEvent(event Event) {
	connection.Events.PushBack(event)
}

func (connection Connection) WithEvent(event Event) Connection {
	connection.Events.PushBack(event)
	return connection
}