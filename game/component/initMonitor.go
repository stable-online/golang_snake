package component

import (
	"github.com/nsf/termbox-go"
)

type MonitorKeyType func(chan int, chan int)

//InitMonitor 初始化监控信息
func InitMonitor() MonitorKeyType {
	return func(keyboard chan int, quit chan int) {

		termbox.SetInputMode(termbox.InputEsc)

		for {
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyArrowLeft:
					keyboard <- LEFT
				case termbox.KeyArrowDown:
					keyboard <- DOWN
				case termbox.KeyArrowRight:
					keyboard <- RIGHT
				case termbox.KeyArrowUp:
					keyboard <- UP
				case termbox.KeyEsc:
					quit <- QUIT
				}
			default:
				keyboard <- 0
			case termbox.EventError:
				panic(ev.Err)
			}
		}
	}
}
