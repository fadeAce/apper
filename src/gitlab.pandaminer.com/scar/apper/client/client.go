package client

import "sync"

type Notifier struct {
	sync.RWMutex
}

func NewNotifier() *Notifier {
	n := &Notifier{}
	return n
}

func (n *Notifier) Notify(key string, data []byte) {
	n.Lock()
	conn.Publish("NOTIFY."+key, data)
	n.Unlock()
}