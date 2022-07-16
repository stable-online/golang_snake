package component

func newGameData() *GameData {
	return &GameData{keyboardChan: make(chan int), quitChan: make(chan int), runtimeChan: make(chan bool, 1), gameOver: false, score: 0, direction: UP}
}
