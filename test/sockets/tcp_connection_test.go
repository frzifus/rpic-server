package sockets_test

import (
	m "github.com/frzifus/rpic-server/messages"
	s "github.com/frzifus/rpic-server/sockets"
	"net"
	"testing"
)

func TestTCPServerAndProto(t *testing.T) {
	tcp := s.NewTCPHandler("127.0.0.1", "4444")
	if tcp == nil {
		t.Fatalf("Couldnt start tcp server")
	}
	go tcp.Listen()
	defer tcp.Dispose()
	c, err := net.Dial("tcp", "127.0.0.1:4444")
	if err != nil {
		t.Fatalf("Connection lost!")
	}
	msg := m.NewCtlMessage(0, 1)
	data, err := m.EncodeCtlMessage(msg)
	if err != nil {
		t.Fatalf("Connection lost! %s", err)
	}
	_, err = c.Write(data)
	if err != nil {
		t.Fatalf("Transmit error! %s", err)
	}
}
