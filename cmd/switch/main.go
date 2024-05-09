// This example show an example of a switch accessory
// which periodically changes it's state between on and off.
package main

import (
	"github.com/LUJUNQUAN/hap"
	"github.com/LUJUNQUAN/hap/accessory"

	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	a := accessory.NewBridge(accessory.Info{
		Name: "bridge",
	})

	a.Id = uint64(1)

	b := accessory.NewSwitch(accessory.Info{
		Name: "Lamp",
	})

	b.Id = uint64(2)

	d := accessory.NewSwitch(accessory.Info{
		Name: "Lamp1",
	})

	d.Id = uint64(3)

	s, err := hap.NewServer(hap.NewFsStore("./db"), a.A, b.A, d.A)
	if err != nil {
		log.Panic(err)
	}

	// Log to console when client (e.g. iOS app) changes the value of the on characteristic
	b.Switch.On.OnValueRemoteUpdate(func(on bool) {
		if on {
			log.Println("Client changed switch to on")
		} else {
			log.Println("Client changed switch to off")
		}
	})

	s.Pin = "34679023"

	// Periodically toggle the switch's on characteristic
	go func() {
		for {
			on := !b.Switch.On.Value()
			if on {
				log.Println("Switch is on")
			} else {
				log.Println("Switch is off")
			}
			b.Switch.On.SetValue(on)
			time.Sleep(1 * time.Second)
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-c
		signal.Stop(c)
		cancel()
	}()

	s.ListenAndServe(ctx)
}
