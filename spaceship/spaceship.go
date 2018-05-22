package spaceship

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/nimrodshn/GoInvaders/utils"
	"golang.org/x/image/colornames"
)

const (
	stepSize = 10
)

// Spaceship struct is a common struct for both player and enemy
type Spaceship struct {
	sprite *pixel.Sprite
	mat    pixel.Matrix
}

// NewMainPlayer Creates a new main player for the game and draws it on a given window.
func NewMainPlayer() (*Spaceship, error) {
	// Initial player location
	initLocation := pixel.V(float64(utils.WindowWidth/2), float64(utils.WindowHeight/10))

	// Load main player sprite.
	sprite, err := utils.LoadSprite("./assets/images/spaceship.png")

	if err != nil {
		return nil, err
	}

	mat := pixel.IM
	mat = mat.Moved(initLocation)

	player := new(Spaceship)
	player.mat = mat
	player.sprite = sprite

	return player, nil
}

// DrawOnScreen draws the spaceship on screen win.
func (player *Spaceship) DrawOnScreen(win *pixelgl.Window) {
	// Set Background Color.
	win.Clear(colornames.Black)
	player.sprite.Draw(win, player.mat)
}

// ListenAndMoveOnKeyStroke Moves player on key strokes
func (player *Spaceship) ListenAndMoveOnKeyStroke(win *pixelgl.Window) {
	var newLocation pixel.Matrix
	switch {
	case win.Pressed(pixelgl.KeyLeft):
		newLocation = player.mat.Moved(pixel.V(-stepSize, 0))
	case win.Pressed(pixelgl.KeyRight):
		newLocation = player.mat.Moved(pixel.V(stepSize, 0))
	case win.Pressed(pixelgl.KeyDown):
		newLocation = player.mat.Moved(pixel.V(0, -stepSize))
	case win.Pressed(pixelgl.KeyUp):
		newLocation = player.mat.Moved(pixel.V(0, stepSize))
	}
	if inBounds(newLocation) {
		player.mat = newLocation
	}
}

func inBounds(mat pixel.Matrix) bool {
	if (mat[4] < utils.WindowWidth && mat[4] > 0) &&
		(mat[5] < utils.WindowHeight && mat[5] > 0) {
		return true
	}
	return false
}
