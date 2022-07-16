package game

import "snake/game/component"

//Start starting snake program
func Start() {
	run(func() { component.NewGameService().Start() })
}
