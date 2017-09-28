package sockets

import (
	"fmt"
	c "github.com/frzifus/rpic-server/camera"
	"net"
	"time"
)

type udpClient struct {
	receiverAddr []string
	connPort     string
	camera       *c.RpicCamera
	interval     time.Duration
	isActive     bool
	// mu        sync.Mutex
}

func NewUDPClient(connHost string, connPort string) *udpClient {
	u := &udpClient{}
	u.receiverAddr = make([]string, 0)
	u.interval = 1 * time.Second
	u.connPort = connPort
	u.isActive = false
	return u
}

func (u *udpClient) SetCamera(cam *c.RpicCamera) {
	u.camera = cam
}

func (u *udpClient) SetInterval(td time.Duration) {
	u.interval = time.Second * td
}

func (u *udpClient) UpdateReceiver(t *tcpHandler) {
	u.receiverAddr = make([]string, len(t.GetAllRemoteAddr()))
	copy(u.receiverAddr, t.GetAllRemoteAddr())
}

func (u *udpClient) IsActive() bool {
	return u.isActive
}

func (u *udpClient) Dispose() {
	u.isActive = false
}

func (u *udpClient) SeedImages() {
	u.isActive = true
	for u.isActive {
		time.Sleep(u.interval)
		msg, err := u.camera.GetProtoMsg()

		if nil != err || len(u.receiverAddr) == 0 {
			continue
		}

		for _, rec := range u.receiverAddr {
			conn, _ := net.Dial("udp", rec+":"+u.connPort)
			defer conn.Close()
			fmt.Println("Send Image:", len(msg), "Bit")
			_, err = conn.Write(msg)
			if err != nil {
				continue
			}
		}
	}
}
