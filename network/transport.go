package network

type NetAddr string

type RPC struct {
	From    NetAddr
	Payload []byte
}

type Transport interface {
	Connect(Transport) error
	Consume() <-chan RPC
	SendMessage(to NetAddr, payload []byte) error
	Addr() NetAddr
}
