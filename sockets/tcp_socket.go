package sockets

import (
	g "github.com/frzifus/rpic-server/gpio"
	m "github.com/frzifus/rpic-server/messages"
	"log"
	"net"
)

type tcpHandler struct {
	listener       net.Listener
	err            error
	msgChan        chan *m.RpicMessage
	shutdown       bool
	tcpConns       []tcpConnection
	connLimit      int
	udpClient      *udpClient
	gpioController *g.GpioController
}

func NewTCPHandler(connHost string, connPort string) *tcpHandler {
	t := &tcpHandler{}
	t.listener, t.err = net.Listen("tcp", connHost+":"+connPort)
	if t.err != nil {
		log.Fatal(t.err)
	}
	t.msgChan = make(chan *m.RpicMessage)
	t.shutdown = false
	t.tcpConns = make([]tcpConnection, 0)
	t.connLimit = 4
	return t
}

func (t *tcpHandler) SetGpioController(g *g.GpioController) {
	t.gpioController = g
	t.gpioController.SetMsgFunc(t.GetMessages)
}

func (t *tcpHandler) IsActive() bool {
	return !t.shutdown
}

func (t *tcpHandler) SetUDPClient(client *udpClient) {
	t.udpClient = client
}

func (t *tcpHandler) RemoveUDPClient() {
	t.udpClient = nil
}

func (t *tcpHandler) notifyUDPClient() {
	t.udpClient.UpdateReceiver(t)
}

func (t *tcpHandler) addTCPConnection(conn net.Conn) {
	if t.connLimit > t.CountConnections() {
		c := NewTCPConnection(conn, t.msgChan, t.removeTCPConnection,
			len(t.tcpConns))
		t.tcpConns = append(t.tcpConns, *c)
		go c.listen()
		t.notifyUDPClient()
	}
}

func (t *tcpHandler) removeTCPConnection(id int) {
	t.tcpConns = append(t.tcpConns[:id], t.tcpConns[id+1:]...)
	t.notifyUDPClient()
}

func (t *tcpHandler) cleanTCPConnections() {
	for i, conn := range t.tcpConns {
		if !conn.IsActive() {
			t.tcpConns = append(t.tcpConns[:i], t.tcpConns[i+1:]...)
			t.notifyUDPClient()
		}
	}
}

func (t *tcpHandler) Listen() {
	for !t.shutdown {
		if conn, err := t.listener.Accept(); err == nil {
			t.addTCPConnection(conn)
		}
	}
}

func (t *tcpHandler) SetConnectionLimit(lim int) {
	t.connLimit = lim
}

func (t *tcpHandler) GetConnectionLimit() int {
	return t.connLimit
}

func (t *tcpHandler) Dispose() {
	t.shutdown = true
}

func (t *tcpHandler) GetMessages() *m.RpicMessage {
	return <-t.msgChan
}

func (t *tcpHandler) GetLastError() error {
	return t.err
}

func (t *tcpHandler) CountConnections() int {
	return len(t.tcpConns)
}

func (t *tcpHandler) GetAllRemoteAddr() []string {
	var res []string
	for _, conn := range t.tcpConns {
		res = append(res, conn.getAddrAsString())
	}
	return res
}

func (t *tcpHandler) ReplyToAll(msg *m.CtlMessage) error {
	for _, conn := range t.tcpConns {
		data, err := m.EncodeCtlMessage(msg)
		if err != nil {
			return err
		}
		_, err = conn.write(data)
		if err != nil {
			return err
		}
	}
	return nil
}
