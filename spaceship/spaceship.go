package spaceship

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/nimrodshn/GoInvaders/bullete"
	"github.com/nimrodshn/GoInvaders/utils"
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

func (player *Spaceship) shoot(win *pixelgl.Window) {
	_, err := bullete.NewBullete(pixel.V(player.mat[4], player.mat[5]))
	if err != nil {
		return
	}
}

func (player *Spaceship) SetMatrix(matrix pixel.Matrix) {
	player.mat = matrix
}

// GetObjectMatrix Return the object matrix containing information needed
// in order to render spaceship.
func (player Spaceship) GetObjectMatrix() pixel.Matrix {
	return player.mat
}

// GetObjectSprite Return the object matrix containing information needed
// in order to render spaceship.
func (player Spaceship) GetObjectSprite() pixel.Sprite {
	return *player.sprite
}
