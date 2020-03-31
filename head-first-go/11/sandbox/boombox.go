// B''H

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/barucharky/go-shenanigans/head-first-go/11/sandbox/gadget"
	"github.com/barucharky/go-shenanigans/head-first-go/11/sandbox/player"
)

func main() {
	// Create tapePlayer
	var tPlayer gadget.TapePlayer
	tPlayer.Batteries = "Duracell"
	tPlayer.Radio = true
	tPlayer.Song = getSongName()

	// Test it out
	player.Process(tPlayer, tPlayer.Song)

	// -- ---------------
	fmt.Println("-----")
	// -- ---------------

	// Create recorder
	var tRecorder gadget.TapeRecorder
	tRecorder.Microphones = 2
	tRecorder.Song = getSongName()

	// Test it out
	player.Process(tRecorder, tRecorder.Song)
	tRecorder.Record()

}

// Get song name
func getSongName() string {

	fmt.Print("Enter song name: ")

	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	var input string
	var err error

	input, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	return input
}
