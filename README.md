event-dispatcher
================

An Event Dispatcher for Go, enforced by Go's first class citizenship of functions.

Experimental. Tests still uncomplete.

Here is a brief example of use:

``` go
package main

import (
	"fmt"
	"github.com/golangit/eventdispatcher/event"
	"github.com/golangit/eventdispatcher/eventdispatcher"
)

type hello struct {
	name string
}

type Hello interface {
	HelloWorld(e event.Event)
	GetName() string
}

func New(name string) Hello {
	return &hello{name: name}
}

func (h *hello) HelloWorld(e event.Event) {
	fmt.Println(h.GetName())
}

func (h *hello) GetName() string {
	return h.name
}

func main() {
	helloWorld := New("Hello World")
	listener := eventdispatcher.Listener{Callable: helloWorld.HelloWorld, Priority: 1}
	ed := eventdispatcher.New()
	event := event.New()
	ed.AddListener("hello", listener)
	ed.Dispatch("hello", event)
}
```
