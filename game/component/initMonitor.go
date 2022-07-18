package component

import (
	"github.com/nsf/termbox-go"
)

type monitorFunType func(chan int, chan int)

//initMonitor 初始化监控
func initMonitor() monitorFunType {
	return func(monitorChan chan int, quit chan int) {

		termbox.SetInputMode(termbox.InputEsc)

		for {
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyArrowLeft:
					monitorChan <- LEFT
				case termbox.KeyArrowDown:
					monitorChan <- DOWN
				case termbox.KeyArrowRight:
					monitorChan <- RIGHT
				case termbox.KeyArrowUp:
					monitorChan <- UP
				case termbox.KeyEsc:
					quit <- QUIT
				}
			case termbox.EventError:
				panic(ev.Err)
			}
		}
	}
}
