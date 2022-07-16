package component

func newGameData() *gameData {
	return &gameData{keyboardChan: make(chan int), quitChan: make(chan int), runtimeChan: make(chan bool, 1), gameOver: false, score: 0, direction: UP}
}
