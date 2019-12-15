// B"H

package gadget

import "fmt"

// -- -------------------------------------
/*
TapePlayer struct
    str   : Batteries string
    method: func (t TapePlayer) Play(song string)
    method: func (t TapePlayer) Stop()
*/
// -- -------------------------------------
type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

// -- -------------------------------------
/*
TapeRecorder struct
    int   : Microphones int
    method: func (t TapeRecorder) Play(song string)
    method: func (t TapeRecorder) Stop()
    method: func (t TapeRecorder) Record()
*/
// -- -------------------------------------
type TapeRecorder struct {
	Microphones int
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopped!")
}
