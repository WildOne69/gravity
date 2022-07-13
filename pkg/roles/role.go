package roles

import (
	"context"

	"beryju.io/gravity/pkg/storage"
	log "github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Role interface {
	Start(ctx context.Context, config []byte) error
	Stop()
}

type Event struct {
	Payload EventPayload
	topic   string
}

func (ev *Event) WithTopic(topic string) *Event {
	ev.topic = topic
	return ev
}

func (ev *Event) String() string {
	return ev.topic
}

type EventPayload struct {
	Data                 map[string]interface{}
	RelatedObjectKey     *storage.Key
	RelatedObjectOptions []clientv3.OpOption
}

func NewEvent(data map[string]interface{}) *Event {
	return &Event{
		Payload: EventPayload{
			Data:                 data,
			RelatedObjectOptions: make([]clientv3.OpOption, 0),
		},
	}
}

type EventHandler func(ev *Event)

type Instance interface {
	KV() *storage.Client
	GetLogger() *log.Entry
	DispatchEvent(topic string, ev *Event)
	AddEventListener(topic string, handler EventHandler)
}
