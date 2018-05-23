package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/nimrodshn/GoInvaders/spaceship"
	"github.com/nimrodshn/GoInvaders/utils"
	"golang.org/x/image/colornames"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	player, gameWindow, err := initializeGame()

	if err != nil {
		panic(err)
	}

	// Main game loop
	for !gameFinished(gameWindow) {
		player.ListenAndMoveOnKeyStroke(gameWindow)
		// After updating the new location we need to rerender to screen
		player.DrawOnScreen(gameWindow)
		gameWindow.Update()
	}
}

func initializeGame() (*spaceship.Spaceship, *pixelgl.Window, error) {
	cfg := pixelgl.WindowConfig{
		Title:  "GoInvaders",
		Bounds: pixel.R(0, 0, utils.WindowWidth, utils.WindowHeight),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// Set Background Color.
	win.Clear(colornames.Black)

	// Create and draw the main player in the game window.
	player, err := spaceship.NewMainPlayer()
	player.DrawOnScreen(win)

	return player, win, err
}

func gameFinished(gameWindow *pixelgl.Window) bool {
	if gameWindow.Closed() {
		return true
	}
	return false
}
