package streamdeck

import (
	"context"
	"sync"

	sdcontext "github.com/FlowingSPDG/streamdeck/context"
)

// Action action instance
type Action struct {
	uuid     string
	handlers eventHandlers
	contexts contexts
}

// map[sring][]EventHandler
type eventHandlers struct {
	m sync.Map
}

// map[string]context.Context
type contexts struct {
	m sync.Map
}

func newAction(uuid string) *Action {
	action := &Action{
		uuid:     uuid,
		handlers: eventHandlers{},
		contexts: contexts{m: sync.Map{}},
	}

	action.RegisterHandler(WillAppear, func(ctx context.Context, client *Client, event Event) error {
		action.addContext(ctx)
		return nil
	})

	action.RegisterHandler(WillDisappear, func(ctx context.Context, client *Client, event Event) error {
		action.removeContext(ctx)
		return nil
	})

	return action
}

// RegisterHandler Register event handler to specified event. handlers can be multiple(append slice)
func (action *Action) RegisterHandler(eventName string, handler EventHandler) {
	handlers, ok := action.handlers.m.Load(eventName)
	if !ok {
		handlers = []EventHandler{}
		action.handlers.m.Store(eventName, handlers)
	}
	hs := handlers.([]EventHandler)
	hs = append(hs, handler)
	action.handlers.m.Store(eventName, hs)
}

// Contexts get contexts
func (action *Action) Contexts() []context.Context {
	cs := make([]context.Context, 0) // 0 length/capacity
	action.contexts.m.Range(func(key, value interface{}) bool {
		v := value.(context.Context)
		cs = append(cs, v)
		return true
	})
	return cs
}

func (action *Action) addContext(ctx context.Context) {
	if sdcontext.Context(ctx) == "" {
		panic("passed non-streamdeck context to addContext")
	}
	action.contexts.m.Store(sdcontext.Context(ctx), ctx)
}

func (action *Action) removeContext(ctx context.Context) {
	if sdcontext.Context(ctx) == "" {
		panic("passed non-streamdeck context to addContext")
	}
	action.contexts.m.Delete(sdcontext.Context(ctx))
}
