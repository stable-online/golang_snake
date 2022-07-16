package game

//Start starting snake program
func Start() {
	run(func() { newGameService().start() })
}
