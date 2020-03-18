// B''H

package player

import "fmt"

// Player is the interface created to play tapes
type Player interface {
	Play(string)
	Stop()
}

// Create tools for interface

// Process logs, plays and stops
func Process(player Player) {

	logPlayInfo(player)
	player.Play("Test Track")
	player.Stop()

}

// func logPlayInfo prints "logging to db"
func logPlayInfo(player Player) {
	fmt.Println("Logging to database...")
}
