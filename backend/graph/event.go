package graph

import "log"

var (
	EventUpdate []func(graphType string, key string)
)

func RegisterEventListener(evType string, f any) {
	switch evType {
	case "update":
		EventUpdate = append(EventUpdate, f.(func(string, string)))
	default:
		log.Println("unknown event type ", evType)
	}
}

func notifyEventUpdate(grType, key string) {
	go func() {
		for _, f := range EventUpdate {
			f(grType, key)
		}
	}()
}
