package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/nimrodshn/GoInvaders/spaceship"
	"github.com/nimrodshn/GoInvaders/utils"
	"golang.org/x/image/colornames"
)

// The Height & Width of the game window
const ()

func main() {
	pixelgl.Run(run)
}

func run() {

	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
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
	player, err := spaceship.NewMainPlayer(win)

	if err != nil {
		panic(err)
	}

	// Main game loop
	for !win.Closed() {
		player.ListenAndMoveOnKeyStroke(win)
		win.Update()
	}
}
