package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/nimrodshn/GoInvaders/userinterface"
	"time"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	ui, err := initializeGame()
	ticker := time.NewTicker(250 * time.Millisecond)

	if err != nil {
		panic(err)
	}

	// Main game loop
	for {
		select {
		case <-ticker.C:
			ui.ListenOnKeyStroke()
			ui.DrawGameOnScreen()
		default:
			if gameFinished(ui) {
				break
			}
		}
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
