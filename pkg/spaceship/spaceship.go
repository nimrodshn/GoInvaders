package spaceship

import (
	"github.com/faiface/pixel"
	"github.com/nimrodshn/GoInvaders/pkg/utils"
)

// Spaceship struct is a common struct for both player and enemy
type Spaceship struct {
	sprite *pixel.Sprite
}

// NewMainPlayer Creates a new main player.
func NewMainPlayer() (*Spaceship, error) {
	// Load main player sprite.
	sprite, err := utils.LoadSprite("./assets/images/spaceship.png")

	if err != nil {
		return nil, err
	}

	player := new(Spaceship)

	player.sprite = sprite

	return player, nil
}

// NewEnemy Creates a new enemy for the game.
func NewEnemy() (*Spaceship, error) {
	// Load main player sprite.
	sprite, err := utils.LoadSprite("./assets/images/invader.png")

	if err != nil {
		return nil, err
	}

	player := new(Spaceship)
	player.sprite = sprite

	return player, nil
}

// GetObjectSprite Return the object sprite containing information needed
// in order to render spaceship.
func (player *Spaceship) GetObjectSprite() *pixel.Sprite {
	return player.sprite
}
