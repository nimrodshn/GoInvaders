package main

import (
	"image"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// The Height & Width of the game window
const (
	WindowWidth  = 1024
	WindowHeight = 768
)

func main() {
	pixelgl.Run(run)
}

func run() {
	initLocation := pixel.V(float64(WindowWidth/2), float64(WindowHeight/10))

	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, WindowWidth, WindowHeight),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)

	if err != nil {
		panic(err)
	}

	pic, err := loadPicture("./assets/images/spaceship.png")
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(pic, pic.Bounds())

	// Set Background Color.
	win.Clear(colornames.Black)

	// Main game loop
	for !win.Closed() {
		mat := pixel.IM
		mat = mat.Moved(initLocation)
		sprite.Draw(win, mat)

		win.Update()
	}
}

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}
