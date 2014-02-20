// eventdispatcher_test.go
package eventdispatcher_test

import (
	"github.com/golangit/eventdispatcher/event"
	. "github.com/golangit/eventdispatcher/eventdispatcher"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestEventDispatcher(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Event Dispatcher Suite")
}
func testfun(e event.Event) { print("hello") }

var _ = Describe("Event Dispatcher", func() {
	It("should create a new instance", func() {
		ed := New()
		Expect(ed).ShouldNot(BeNil())
	})

	var (
		listener1 = Listener{Callable: testfun, Priority: 0}
		listener2 = Listener{Callable: testfun, Priority: 1}

		listeners = map[string]Listener{
			"foo": listener1,
			"bar": listener2,
		}
	)
	Context("When I register a listener", func() {
		It("Should be available among eventdispatcher listeners", func() {
			ed := New()
			for k, v := range listeners {
				ed.AddListener(k, v)
			}
			Expect(ed.GetListeners("foo")[0].Callable).ShouldNot(BeNil())
		})
	})
})
