// B"H

package main

import "github.com/barucharky/go-shenanigans/head-first-go/11/11b/gadget"

// -- -------------------------------------
type Player interface {
	Play(string)
	Stop()
}

// -- -------------------------------------
func tryOutTapePlayer(t gadget.TapePlayer) {
	t.Play("Hashem is here!")
	t.Stop()
}

// -- -------------------------------------
func TryOutGadget(player Player) {

	player.Play("Test Track")
	player.Stop()

	// -- ---------------------------------
	// 99% if time you won't need this:
	// type assertion:
	recorder, ok := player.(gadget.TapeRecorder)

	if ok {
		recorder.Record()
	}

}

// -- -------------------------------------
func main() {

	// -- ---------------------------------
	var tPlayer gadget.TapePlayer
	tPlayer.Batteries = "Duracell"

	tryOutTapePlayer(tPlayer)

	// -- ---------------------------------
	var tRecorder gadget.TapeRecorder
	tRecorder.Microphones = 2

	// -- ---------------------------------
	TryOutGadget(tPlayer)
	TryOutGadget(tRecorder)

}
