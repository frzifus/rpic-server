package main

import (
	"fmt"
	c "github.com/frzifus/rpic-server/camera"
	g "github.com/frzifus/rpic-server/gpio"
	s "github.com/frzifus/rpic-server/sockets"
	"os"
	"os/signal"
	"syscall"
)

const (
	cameraPath = "/dev/video0"
	addr       = "0.0.0.0"
	portIn     = "4445"
	portOut    = "4444"
)

func init() {
	fmt.Println("Start rpic-rerver!")
}

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)

	cam := c.NewRpicCamera(cameraPath)
	gpio := g.NewGpioController()
	defer gpio.Dispose()
	tcp := s.NewTCPHandler(addr, "4444")
	defer tcp.Dispose()
	udp := s.NewUDPClient(addr, "4445")
	defer udp.Dispose()

	tcp.SetUDPClient(udp)
	tcp.SetGpioController(gpio)
	udp.SetCamera(cam)

	go tcp.Listen()
	go udp.SeedImages()
	go gpio.Start()

	go func() {
		for {
			sig := <-sigs
			if sig == syscall.SIGINT || sig == syscall.SIGTERM {

				fmt.Println()
				fmt.Println("signal =>", sig)
				fmt.Println("shutdown rpic-server!")
				done <- true
				return
			} else if sig == syscall.SIGUSR1 {
				if tcp.IsActive() {
					fmt.Println("TCP - up!")
				}
				if udp.IsActive() {
					fmt.Println("UDP - up!")
				}
				fmt.Println("Connected:", tcp.GetAllRemoteAddr())
			}
		}
	}()
	<-done
}
