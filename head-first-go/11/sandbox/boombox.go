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

	// Make a TapePlayer
	var tPlayer gadget.TapePlayer
	tPlayer.Batteries = "Duracell"
	tPlayer.Radio = false
	tPlayer.Song = getSong()

	// Test the TapePlayer
	player.Process(tPlayer)

	// -- -------------------------------
	fmt.Println("----------")
	// -- -------------------------------

	// Make a TapeRecorder
	var tRecorder gadget.TapeRecorder
	tRecorder.Microphones = 2
	tRecorder.Song = getSong()

	// Test the TapeRecorder
	player.Process(tRecorder)
	tRecorder.Record()
}

func getSong() string {

	fmt.Println("Enter song name:")

	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var input string
	var err error

	input, err = reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	return input
}
