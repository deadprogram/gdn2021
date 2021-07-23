// +build baremetal,!circuitplay_bluefruit

package main

import "machine"

var (
	on  = false
	led = machine.LED
)

func init() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
}

func showPulse() {
	if on {
		led.Low()
	} else {
		led.High()
	}
	on = !on
}
