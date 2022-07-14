package main

import "snake_game/game"

func main() {

	//to create the game starter
	GameStarter := game.Games(game.Screen(game.InitSnake(), game.InitFood(), game.InitMove()), game.MonitorKeyboard())

	//starting
	GameStarter()
}
