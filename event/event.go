package event

type event struct {
	dispatcher interface{}
	name       string
}

type Event interface {
	SetDispatcher(dispatcher interface{})
	SetName(name string)
	IsPropagationStopped() bool
}

func New() Event {
	return &event{}
}

func (e *event) SetDispatcher(dispatcher interface{}) {
	e.dispatcher = dispatcher
}

func (e *event) SetName(name string) {
	e.name = name
}

func (e *event) IsPropagationStopped() bool {
	return false
}
