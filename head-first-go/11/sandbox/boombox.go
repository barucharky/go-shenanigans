// B''H

package main

import (
	"fmt"

	"github.com/barucharky/go-shenanigans/head-first-go/11/sandbox/gadget"
	"github.com/barucharky/go-shenanigans/head-first-go/11/sandbox/player"
)

func main() {
	// Create tapePlayer
	var tPlayer gadget.TapePlayer
	tPlayer.Batteries = "Duracell"

	// Test it out
	player.Process(tPlayer)

	// -- ---------------
	fmt.Println("-----")
	// -- ---------------

	// Create recorder
	var tRecorder gadget.TapeRecorder
	tRecorder.Microphones = 2

	// Test it out
	player.Process(tRecorder)
	tRecorder.Record()

}
