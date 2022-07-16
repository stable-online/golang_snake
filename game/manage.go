package game

import . "snake/game/component"

//Start starting snake program
func Start() {
	Run(func() { NewGameService().Start() })
}
