// +build !baremetal

package main

import "time"

func showPulse() {
	println("tick", time.Now().Format("04:05.000"))
}
