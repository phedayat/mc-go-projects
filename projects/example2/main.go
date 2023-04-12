package main

import "machine"
import "time"

func main(){
	led := machine.LED
	pinConfig := machine.PinConfig{Mode: machine.PinOutput}
	led.Configure(pinConfig)
	sleepTime := time.Millisecond*500
	for {
		led.Low()
		time.Sleep(sleepTime)

		led.High()
		time.Sleep(sleepTime)
	}
}

