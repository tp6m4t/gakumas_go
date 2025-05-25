package field

import "LocalProject/pkg/tool/event"

func (f *Field) Subscribe(name string, value func(interface{})) {
	f.eventBus.Subscribe(name, (func(eventData event.Eventdata) { value(eventData) }))
}
