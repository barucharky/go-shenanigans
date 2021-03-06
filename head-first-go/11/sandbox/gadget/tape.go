package gadget

import "fmt"

// Tape Player struct and methods

type TapePlayer struct {
	Batteries string
	Radio     bool
	Song      string
}

func (t TapePlayer) Play() {
	fmt.Print("Now Playing: ", t.Song)
	fmt.Println("On a Tape Player")
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped -- Tape Player")
}

// Tape Recorder struct and methods

type TapeRecorder struct {
	Microphones int64
	Song        string
}

func (t TapeRecorder) Play() {
	fmt.Print("Now Playing: ", t.Song)
	fmt.Println("On a Tape Recorder")
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopped -- Tape Recorder")
}

func (t TapeRecorder) Record() {
	fmt.Println("Recording...")
}
