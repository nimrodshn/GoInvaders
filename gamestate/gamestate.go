package gamestate

import (
	"github.com/nimrodshn/GoInvaders/gameobject"
	"github.com/nimrodshn/GoInvaders/bullete"
	"github.com/nimrodshn/GoInvaders/spaceship"
	"github.com/faiface/pixel"
)

// GameState holds the current game state
type GameState struct {
	mainPlayer *spaceship.Spaceship
	enemies    []*spaceship.Spaceship
	bullets    []*bullete.Bullete
}

// NewGameState Creates  a new GameState for game initialization
func NewGameState() (state *GameState, err error) {
	state = new(GameState)
	player, err := spaceship.NewMainPlayer()
	state.mainPlayer = player
	return state, err
}

// GetGameObjects returns a snapshot of the current entities in the game.
func (state *GameState) GetGameObjects() []gameobject.GameObject {
	objects := make([]gameobject.GameObject,0)
	// cannot append []T to and interface (see https://golang.org/doc/faq#convert_slice_of_interface).
	for _,enemy := range state.enemies {
		objects = append(objects,enemy)
	}
	for _,bullet := range state.bullets {
		objects = append(objects, bullet)
	}
	objects = append(objects, state.mainPlayer)
	return objects
}

func (state *GameState) GetMainPlayer() *spaceship.Spaceship {
	return state.mainPlayer
}

func (state *GameState) ChangePlayerState(matrix pixel.Matrix) {
	state.mainPlayer.SetMatrix(matrix)
}
