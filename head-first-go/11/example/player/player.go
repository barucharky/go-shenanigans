// B''H

package player

import "fmt"

type Player interface {
	Play(string)
	Stop()
}

func Process(player Player) {

	logPlayInfo(player)
	player.Play("Test Track")
	player.Stop()

}

func logPlayInfo(player Player) {
	fmt.Println("Logging to database...")
}
