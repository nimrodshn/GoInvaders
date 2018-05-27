package shot

import (
	"github.com/faiface/pixel"
	"github.com/nimrodshn/GoInvaders/utils"
)

// Shot The class wrapping the logic of shots either from enemy or main player.
type Shot struct {
	sprite *pixel.Sprite
	mat    pixel.Matrix
}

// NewShot Creates new Shot and render it to screen
func NewShot(spaceshipLoaction pixel.Vec) (*Shot, error) {
	// Load main player sprite.
	sprite, err := utils.LoadSprite("./assets/images/shot.png")

	if err != nil {
		return nil, err
	}

	mat := pixel.IM
	mat = mat.Moved(spaceshipLoaction.Add(pixel.V(0, utils.StepSize)))

	shot := new(Shot)
	shot.mat = mat
	shot.sprite = sprite

	return shot, nil
}

// GetObjectMatrix Return the object matrix containing information needed
// in order to render spaceship.
func (shot *Shot) GetObjectMatrix() pixel.Matrix {
	return shot.mat
}

// GetObjectSprite Return the object matrix containing information needed
// in order to render spaceship.
func (shot *Shot) GetObjectSprite() pixel.Sprite {
	return *shot.sprite
}
