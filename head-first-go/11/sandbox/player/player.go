package player

import "fmt"

type Player interface {
	Play()
	Stop()
}

func Process(plyr Player) {
	logPlayerInfo(plyr)
	plyr.Play()
	plyr.Stop()
}

func logPlayerInfo(plyr Player) {
	fmt.Println("Logging to database...")
}
