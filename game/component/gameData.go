package component

func NewGameData() *GameData {
	return &GameData{keyboardChan: make(chan int), quitChan: make(chan int), runtimeChan: make(chan bool, 1)}
}
