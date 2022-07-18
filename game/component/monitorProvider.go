package component

//main
const (
	RIGHT = 1 + iota
	LEFT
	UP
	DOWN
	QUIT
)

type monitorProvider struct {
	Monitor monitorFunType
}

func newMonitorApp() *monitorProvider {
	return &monitorProvider{Monitor: initMonitor()}
}

func (s *monitorProvider) start(GameData *gameData) {
	s.Monitor(GameData.keyboardChan, GameData.quitChan)
}
