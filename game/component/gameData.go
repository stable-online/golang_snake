package component

func NewGameData() *GameData {
	return &GameData{keyboardChan: make(chan int)}
}
