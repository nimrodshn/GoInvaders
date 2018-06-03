package bullete

import (
	"github.com/faiface/pixel"
	"github.com/nimrodshn/GoInvaders/utils"
)

// Bullete The class wrapping the logic of shots either from enemy or main player.
type Bullete struct {
	sprite *pixel.Sprite
	mat    pixel.Matrix
}

// NewBullete Creates new Shot and render it to screen
func NewBullete(spaceshipLoaction pixel.Vec) (*Bullete, error) {
	// Load main player sprite.
	sprite, err := utils.LoadSprite("./assets/images/bullete.png")

	if err != nil {
		return nil, err
	}

	mat := pixel.IM
	mat = mat.Moved(spaceshipLoaction.Add(pixel.V(0, utils.StepSize*5)))

	bullete := new(Bullete)
	bullete.mat = mat
	bullete.sprite = sprite

	return bullete, nil
}

// GetObjectMatrix Return the object matrix containing information needed
// in order to render spaceship.
func (b *Bullete) GetObjectMatrix() pixel.Matrix {
	return b.mat
}

// GetObjectSprite Return the object matrix containing information needed
// in order to render spaceship.
func (b *Bullete) GetObjectSprite() *pixel.Sprite {
	return b.sprite
}

// SetMatrix allows to set bullete matrix.
func (b *Bullete) SetMatrix(matrix pixel.Matrix) {
	b.mat = matrix
}
