package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/nimrodshn/GoInvaders/pkg/userinterface"
	"time"
)

var ticker = time.NewTicker(25 * time.Millisecond)

func main() {
	pixelgl.Run(run)
}

func run() {
	ui, err := initializeGame()

	if err != nil {
		panic(err)
	}

	// Main game loop
	for !gameFinished(ui) {
		<-ticker.C
		ui.ListenOnKeyStroke()
		ui.DrawGameOnScreen()
	}
}

func initializeGame() (userinterface.UserInterface, error) {
	ui, err := userinterface.NewUserInterface()
	if err != nil {
		panic(err)
	}
	return ui, err
}

func gameFinished(ui userinterface.UserInterface) bool {
	return ui.WindowClosed()
}
