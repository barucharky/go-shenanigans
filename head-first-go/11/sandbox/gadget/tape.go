// B''H

package gadget

import "fmt"

// Create tapePlayer struct and methods
type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped -- Player")
}

// Create tapeRecorder struct and methods
type TapeRecorder struct {
	Microphones int64
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopped -- Recorder")
}

func (t TapeRecorder) Record() {
	fmt.Println("Recording...")
}
