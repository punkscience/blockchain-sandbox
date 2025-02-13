package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocalTransport_Connect(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	assert.Equal(t, tra.(*LocalTransport).peers[trb.(*LocalTransport).addr], trb)
	assert.Equal(t, trb.(*LocalTransport).peers[tra.(*LocalTransport).addr], tra)
}

func TestSendMessage(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	msg := []byte("hello world")
	assert.Nil(t, tra.SendMessage(trb.(*LocalTransport).addr, msg))

	rpc := <-trb.Consume()

	assert.Equal(t, rpc.Payload, msg)
	assert.Equal(t, rpc.From, tra.(*LocalTransport).addr)
}
