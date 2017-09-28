package rpio

import ()

const (
	bcm2835Base = 0x20000000
	piGPIOBase  = bcm2835Base + 0x200000
	memLength   = 4096
)

func Init() {}

type Pin uint8

type gpio struct{}

func NewGpio() *gpio {
	return &gpio{}
}
