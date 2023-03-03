package game

import (
	"fmt"
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
	}
}
