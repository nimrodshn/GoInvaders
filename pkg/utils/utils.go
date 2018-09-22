package utils

import (
	"image"
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
)

const (
	// WindowWidth is the game windoe width
	WindowWidth = 1024
	// WindowHeight is the game window height
	WindowHeight = 768
	// StepSize is the speed of the moving spaceship / it's bullete
	StepSize = 10
)

// LoadSprite loads the image in path and returns a sprite with that image borders.
func LoadSprite(path string) (*pixel.Sprite, error) {
	pic, err := loadPicture(path)
	if err != nil {
		return nil, err
	}

	sprite := pixel.NewSprite(pic, pic.Bounds())
	return sprite, nil

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
