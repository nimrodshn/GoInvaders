package bullete

import (
	"github.com/faiface/pixel"
	"github.com/nimrodshn/GoInvaders/utils"
)

// Bullete The class wrapping the logic of shots either from enemy or main player.
type Bullete struct {
	sprite *pixel.Sprite
}

// NewBullete Creates new Shot and render it to screen
func NewBullete() (*Bullete, error) {
	// Load main player sprite.
	sprite, err := utils.LoadSprite("./assets/images/bullete.png")

	if err != nil {
		return nil, err
	}

	bullete := new(Bullete)
	bullete.sprite = sprite

	return bullete, nil
}

// GetObjectSprite Return the object sprite containing information needed
// in order to render spaceship.
func (b *Bullete) GetObjectSprite() *pixel.Sprite {
	return b.sprite
}
