// B''H

package gadget

import "fmt"

// Create TapePlayer struct and methods
type TapePlayer struct {
	Batteries string
	Radio     bool
	Song      string
}

func (t TapePlayer) Play(song string) {
	fmt.Print("Now playing:", song)
	fmt.Println("On Tape Player")
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped -- Tape Player")
}

// Create TapeRecorder struct and methods
type TapeRecorder struct {
	Microphones int64
	Song        string
}

func (t TapeRecorder) Play(song string) {
	fmt.Print("Now playing:", song)
	fmt.Println("On Tape Recorder")
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopped -- Tape Recorder")
}

func (t TapeRecorder) Record() {
	fmt.Println("Recording...")
}
