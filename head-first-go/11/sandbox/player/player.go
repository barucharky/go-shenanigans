// B''H

package player

import "fmt"

// Create interface
type Player interface {
	Play(string)
	Stop()
}

// Create tools for interface

// func Process logs, plays and stops
func Process(player Player) {
	logPlayInfo(player)
	player.Play("Test Track")
	player.Stop()
}

// func logPlayInfo prints "logging to db"
func logPlayInfo(player Player) {
	fmt.Println("Logging to database...")
}
