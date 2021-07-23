// +build baremetal,circuitplay_bluefruit

package main

import (
	"image/color"
	"machine"

	"tinygo.org/x/drivers/ws2812"
)

var ledColor = [3]byte{0x00, 0xff, 0x00} // start out with red
var leds [10]color.RGBA

var neo machine.Pin = machine.NEOPIXELS
var ws ws2812.Device
var rg bool

func init() {
	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ws = ws2812.New(neo)

}

func showPulse() {
	rg = !rg

	for i := range leds {
		rg = !rg
		if rg {
			leds[i] = color.RGBA{R: ledColor[0], G: ledColor[1], B: ledColor[2]}
		} else {
			leds[i] = color.RGBA{R: 0x00, G: 0x00, B: 0x00}
		}
	}

	ws.WriteColors(leds[:])
}

func clearLEDS() {
	for i := range leds {
		leds[i] = color.RGBA{R: 0x00, G: 0x00, B: 0x00}
	}

	ws.WriteColors(leds[:])
}
