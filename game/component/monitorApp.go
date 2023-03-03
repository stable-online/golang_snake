package component

import "github.com/nsf/termbox-go"

//main
const (
	RIGHT = 1 + iota
	LEFT
	UP
	DOWN
	QUIT
)

// Monitor
//
// @Description:
type Monitor interface {
	// start
	//
	// @Description:  启动监控
	// @return *monitorApp
	start() *monitorApp
}

// monitorApp
// @Description:
type monitorApp struct {
	move chan int
	quit chan int
}

//newMonitorApp 初始化监控应用
func newMonitorApp() Monitor {
	return &monitorApp{make(chan int), make(chan int)}
}

//handle 启动监听
func (s *monitorApp) start() *monitorApp {
	//异步监控
	go s.initMonitor()
	return s
}

// initMonitor
//
// @Description:
// @receiver s
// @param monitorChan
// @param quit
func (s *monitorApp) initMonitor() {
	termbox.SetInputMode(termbox.InputEsc)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowLeft:
				s.move <- LEFT
			case termbox.KeyArrowDown:
				s.move <- DOWN
			case termbox.KeyArrowRight:
				s.move <- RIGHT
			case termbox.KeyArrowUp:
				s.move <- UP
			case termbox.KeyEsc:
				s.quit <- QUIT
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
