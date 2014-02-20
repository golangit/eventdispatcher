package eventdispatcher

import (
	"github.com/golangit/eventdispatcher/event"
	"sort"
)

type Listener struct {
	Callable func(e event.Event)
	Priority int
}

type ListenersByPriority []Listener

func (l ListenersByPriority) Len() int           { return len(l) }
func (l ListenersByPriority) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ListenersByPriority) Less(i, j int) bool { return l[i].Priority < l[j].Priority }

type EventSubscriber interface {
	GetSubscribedEvents() map[string]Listener
}

type EventDispatcher interface {
	Dispatch(eventName string, event event.Event) event.Event
	AddListener(eventName string, listener Listener)
	AddSubscriber(subscriber EventSubscriber)
	GetListeners(eventName string) []Listener
}

type eventdispatcher struct {
	listeners map[string][]Listener
	sorted    map[string][]Listener
}

func New() EventDispatcher {
	return &eventdispatcher{
		listeners: make(map[string][]Listener),
		sorted:    make(map[string][]Listener),
	}
}

func (ed *eventdispatcher) Dispatch(eventName string, event event.Event) event.Event {
	event.SetDispatcher(ed)
	event.SetName(eventName)
	if ed.listeners[eventName] == nil {
		return event
	}
	ed.DoDispatch(ed.GetListeners(eventName), eventName, event)
	return event
}

func (ed *eventdispatcher) AddListener(eventName string, listener Listener) {
	ed.listeners[eventName] = append(ed.listeners[eventName], listener)
}

func (ed *eventdispatcher) AddSubscriber(subscriber EventSubscriber) {
	for eventName, listener := range subscriber.GetSubscribedEvents() {
		ed.AddListener(eventName, listener)
	}
	return
}

func (ed *eventdispatcher) DoDispatch(listeners []Listener, eventName string, event event.Event) {
	for k := range listeners {
		listeners[k].Callable(event)
		if event.IsPropagationStopped() {
			break
		}
	}

	return
}

func (ed *eventdispatcher) GetListeners(eventName string) []Listener {
	if nil == ed.sorted[eventName] {
		ed.SortListeners(eventName)
	}
	return ed.sorted[eventName]
}

func (ed *eventdispatcher) SortListeners(eventName string) {
	sort.Sort(ListenersByPriority(ed.listeners[eventName]))
	if ed.sorted[eventName] == nil {
		ed.sorted[eventName] = make([]Listener, 1)
	}
	ed.sorted[eventName] = ed.listeners[eventName]
}

func (ed *eventdispatcher) HasListeners(eventName string) bool {
	return 0 < len(ed.GetListeners(eventName))
}
