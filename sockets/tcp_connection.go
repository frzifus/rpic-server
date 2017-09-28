package sockets

import (
	m "github.com/frzifus/rpic-server/messages"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

type tcpConnection struct {
	id       int
	conn     net.Conn
	isActive bool
	msgChan  chan *m.RpicMessage
	dispose  func(int)
	timeout  time.Duration
}

func NewTCPConnection(c net.Conn, msg chan *m.RpicMessage,
	dp func(int), id int) *tcpConnection {
	t := &tcpConnection{}
	t.isActive = false
	t.conn = c
	t.dispose = dp
	t.id = id
	t.msgChan = msg
	t.timeout = 300
	return t
}

func (t *tcpConnection) IsActive() bool {
	return t.isActive
}

func (t *tcpConnection) write(b []byte) (int, error) {
	t.conn.Write(b)
	return t.conn.Write(b)
}

func (t *tcpConnection) getAddrAsString() string {
	return strings.Split(t.conn.RemoteAddr().String(), ":")[0]
}

func (t *tcpConnection) listen() {
	defer func() {
		t.isActive = false
		t.conn.Close()
		t.dispose(t.id)
	}()

	t.isActive = true
	data := make([]byte, 4096)

	for t.isActive {
		err := t.conn.SetReadDeadline(time.Now().Add(t.timeout * time.Second))
		if err != nil {
			log.Println("SetReadDeadline failed:", err)
			return
		}
		// Read the data waiting on the connection and put it in the data buffer
		n, err := t.conn.Read(data[:])
		if err == io.EOF || n == 0 {
			t.isActive = false
			continue
		}
		// Decoding Protobuf message
		protodata, err := m.DecodeCtlMessage(data, n)
		if err != nil {
			continue
		}
		// Push the protobuf message into a channel
		t.msgChan <- &m.RpicMessage{Wrapped: protodata}
	}
}
