package streamdeck

import (
	"context"
	"encoding/json"

	sdcontext "github.com/FlowingSPDG/streamdeck/context"
)

type Event struct {
	Action     string          `json:"action,omitempty"`
	Event      string          `json:"event,omitempty"`
	UUID       string          `json:"uuid,omitempty"`
	Context    string          `json:"context,omitempty"`
	Device     string          `json:"device,omitempty"`
	DeviceInfo DeviceInfo      `json:"deviceInfo,omitempty"`
	Payload    json.RawMessage `json:"payload,omitempty"`
}

type DeviceInfo struct {
	DeviceName string     `json:"deviceName,omitempty"`
	Type       DeviceType `json:"type,omitempty"`
	Size       DeviceSize `json:"size,omitempty"`
}

type DeviceSize struct {
	Columns int `json:"columns,omitempty"`
	Rows    int `json:"rows,omitempty"`
}

type DeviceType int

const (
	// StreamDeck 15LCD key streamdeck
	StreamDeck DeviceType = 0
	// StreamDeckMini 6LCD key streamdeck mini
	StreamDeckMini DeviceType = 1
	// StreamDeckXL 32LCD key streamdeck mini
	StreamDeckXL DeviceType = 2
	// StreamDeckMobile Streamdeck mobile app on iphone
	StreamDeckMobile DeviceType = 3
)

// NewEvent Generate new event data
func NewEvent(ctx context.Context, name string, payload interface{}) Event {
	p, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	return Event{
		Event:   name,
		Action:  sdcontext.Action(ctx),
		Context: sdcontext.Context(ctx),
		Device:  sdcontext.Device(ctx),
		Payload: p,
	}
}

func NewEventSendToPropertyInspector(ctx context.Context, payload interface{}) Event {
	p, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	return Event{
		Action:  sdcontext.Action(ctx),
		Event:   SendToPropertyInspector,
		Context: sdcontext.Context(ctx),
		Payload: p,
	}
}

func NewEventSetSettings(ctx context.Context, payload interface{}) Event {
	p, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	return Event{
		Event:   SetSettings,
		Context: sdcontext.Context(ctx),
		Payload: p,
	}
}

func NewEventGetSettings(ctx context.Context) Event {
	return Event{
		Event:   GetSettings,
		Context: sdcontext.Context(ctx),
	}
}

func NewEventSetGlobalSettings(ctx context.Context, payload interface{}) Event {
	p, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	return Event{
		Event:   GetSettings,
		Context: sdcontext.Context(ctx),
		Payload: p,
	}
}

func NewEventGetGlobalSettings(ctx context.Context) Event {
	return Event{
		Event:   GetSettings,
		Context: sdcontext.Context(ctx),
	}
}
