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

type eventHandlers struct {
	mutex sync.Mutex
	m     map[string][]EventHandler
}

// map[string]context.Context
type contexts struct {
	m sync.Map
}

func newAction(uuid string) *Action {
	action := &Action{
		uuid: uuid,
		handlers: eventHandlers{
			mutex: sync.Mutex{},
			m:     map[string][]EventHandler{},
		},
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
	action.handlers.mutex.Lock()
	defer action.handlers.mutex.Unlock()

	_, ok := action.handlers.m[eventName]
	if !ok {
		action.handlers.m[eventName] = []EventHandler{}
	}
	action.handlers.m[eventName] = append(action.handlers.m[eventName], handler)
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
