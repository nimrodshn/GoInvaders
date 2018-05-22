package spaceship

import (
	"fmt"

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
func NewMainPlayer(win *pixelgl.Window) (*Spaceship, error) {
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
	sprite.Draw(win, mat)

	return player, nil
}

// ListenAndMoveOnKeyStroke Moves player on key strokes and draw new location on screen
func (player *Spaceship) ListenAndMoveOnKeyStroke(win *pixelgl.Window) {
	fmt.Println(player.mat)
	if inBounds(player) {
		if win.Pressed(pixelgl.KeyLeft) {
			player.mat = player.mat.Moved(pixel.V(-stepSize, 0))
		}
		if win.Pressed(pixelgl.KeyRight) {
			player.mat = player.mat.Moved(pixel.V(stepSize, 0))
		}
		if win.Pressed(pixelgl.KeyDown) {
			player.mat = player.mat.Moved(pixel.V(0, -stepSize))
		}
		if win.Pressed(pixelgl.KeyUp) {
			player.mat = player.mat.Moved(pixel.V(0, stepSize))
		}
		win.Clear(colornames.Black)
		player.sprite.Draw(win, player.mat)
	}
}

func inBounds(spaceship *Spaceship) bool {
	if (spaceship.mat[4] < utils.WindowWidth && spaceship.mat[4] > 0) &&
		(spaceship.mat[5] < utils.WindowHeight && spaceship.mat[5] > 0) {
		return true
	}
	return false
}
