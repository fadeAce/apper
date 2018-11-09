package test

import (
	"fmt"
	"github.com/nats-io/go-nats"
	"testing"
	"time"
)

func Test_nats_clientA(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)
	// Simple Publisher
	nc.Publish("foo", []byte("Hello World"))
	time.Sleep(2 * time.Second)
	// Close connection
	nc.Close()
}

func Test_nats_clientB(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)
	ch := make(chan interface{})
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
		// Close connection
		nc.Close()
		close(ch)
	})
	<-ch
}
