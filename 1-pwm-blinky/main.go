package main

import (
	"fmt"
	"machine"
	"time"
)

var period uint64 = 1e9 / 500

func main() {
	// This program is specific to the Raspberry Pi Pico.
	pin := machine.LED
	pwm := machine.PWM4 // Pin 25 (LED on pico) corresponding to PWM4.

	pwm.Configure(machine.PWMConfig{
		Period: period,
	})

	ch, err := pwm.Channel(pin)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for {
		for i := 1; i < 255; i++ {
			// This performs a fade-out blink
			pwm.Set(ch, pwm.Top()/uint32(i))
			time.Sleep(time.Microsecond * 5)
		}
	}

}
