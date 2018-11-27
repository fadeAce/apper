package types

import (
	"sync"
	"github.com/nats-io/go-nats"
)

type Notifier struct {
	sync.RWMutex
	Conn *nats.Conn
}

func NewNotifier() *Notifier {
	n := &Notifier{}
	return n
}

func (n *Notifier) Notify(key string, data []byte) {
	n.Lock()
	n.Conn.Publish("NOTIFY."+key, data)
	n.Unlock()
}