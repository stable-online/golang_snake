package component

//gameData game 数据
type gameData struct {
	keyboardChan chan int
	quitChan     chan int
	runtimeChan  chan bool
	gameOver     bool
	score        int
	direction    int
}

func newGameData() *gameData {
	return &gameData{keyboardChan: make(chan int), quitChan: make(chan int), runtimeChan: make(chan bool, 1), gameOver: false, score: 0, direction: UP}
}
