// B''H

package player

import "fmt"

// Player is the interface for tapes
type Player interface {
	Play()
	Stop()
}

// -- -------------------------------
// Create tools for interface

// Process logs, plays and stops
func Process(plyr Player) {
	logPlayInfo(plyr)
	plyr.Play()
	plyr.Stop()
}

// logPlayInfo prints "logging to db"
func logPlayInfo(plyr Player) {
	fmt.Println("Logging to database...")
}
