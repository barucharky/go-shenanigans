// B''H

package player

import "fmt"

// Player is the interface for tapes
type Player interface {
	Play(string)
	Stop()
}

// -- -------------------------------
// Create tools for interface

// Process logs, plays and stops
func Process(plyr Player, song string) {
	logPlayInfo(plyr)
	plyr.Play(song)
	plyr.Stop()
}

// logPlayInfo prints "logging to db"
func logPlayInfo(plyr Player) {
	fmt.Println("Logging to database...")
}
