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
	fmt.Println("Playing:", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped -- Player")
}

// Create TapeRecorder struct and methods
type TapeRecorder struct {
	Microphones int64
	Song        string
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Now Playing:", song)
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopped -- Recorder")
}

func (t TapeRecorder) Record() {
	fmt.Println("Recording...")
}
