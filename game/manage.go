package game

import (
	"bytes"
	"fmt"
	"runtime"
	"snake/game/component"
)

// Start starting snake program
//
// @Description:
func Start() {

	//捕获异常
	defer recoverFailed()

	//运行游戏
	component.NewGameService().Start()
}

// recoverFailed
//
// @Description:
func recoverFailed() {
	if pM := recover(); pM != nil {
		fmt.Println(pM)
		fmt.Println(PanicTrace())
	}
}

func PanicTrace() string {
	buf := new(bytes.Buffer)
	for i := 1; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
	}
	return buf.String()
}
