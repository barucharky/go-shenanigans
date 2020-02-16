package main

import (
	"fmt"

	"github.com/barucharky/go-shenanigans/head-first-go/11/sandbox/gadget"
	"github.com/barucharky/go-shenanigans/head-first-go/11/sandbox/player"
)

func main() {

	// make a tapePlayer with batteries
	var tPlayer gadget.TapePlayer
	tPlayer.Batteries = "Duracell"

	// try out the tapePlayer
	player.Process(tPlayer)

	// -- -----------------------------
	fmt.Println("-----")
	// -- -----------------------------

	// Make a tapeRecorder with 2 mic's
	var tRecorder gadget.TapeRecorder
	tRecorder.Microphones = 2

	player.Process(tRecorder)
	tRecorder.Record()
	// -- -----------------------------

	// Try out both with interface

}
