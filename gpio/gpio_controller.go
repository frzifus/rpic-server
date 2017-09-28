package rpio

import (
	"fmt"
	m "github.com/frzifus/rpic-server/messages"
)

type GpioController struct {
	isActive bool
	getMsg   func() *m.RpicMessage
}

func NewGpioController() *GpioController {
	return &GpioController{isActive: true}
}

func (g *GpioController) Start() {
	for g.isActive {
		if g.getMsg == nil {
			continue
		}
		message := g.getMsg()
		switch message.Wrapped.(type) {
		case *m.CtlMessage:
			var msg *m.CtlMessage
			msg = message.Wrapped.(*m.CtlMessage)
			fmt.Println(msg.GetCtlType(), msg.GetDirection(), msg.GetValue())
		default:
			fmt.Println("unknown message type")
		}
	}
}

func (g *GpioController) SetMsgFunc(fn func() *m.RpicMessage) {
	g.getMsg = fn
}

func (g *GpioController) Dispose() {
	g.isActive = false
}
