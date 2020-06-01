package player

import "fmt"

type Player interface {
	Play(string)
	Stop()
}

func Process(plyr Player, song string) {
	logPlayerInfo(plyr)
	plyr.Play(song)
	plyr.Stop()
}

func logPlayerInfo(plyr Player) {
	fmt.Println("Logging to database...")
}
