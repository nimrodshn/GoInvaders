package gamestate

import (
	"github.com/nimrodshn/GoInvaders/utils"
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

const (
	// PlayerMovedLeft constant indicating where the player moved
	PlayerMovedLeft = 1
	// PlayerMovedRight constant indicating where the player moved
	PlayerMovedRight = 2
	// PlayerMovedUp constant indicating where the player moved
	PlayerMovedUp = 3
	// PlayerMovedDown constant indicating where the player moved
	PlayerMovedDown = 4
	// PlayerShotBullet constant indicating the player shot a bullet
	PlayerShotBullet = 5
)

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

// ChangeState changes the player state according to the input given by ui.
func (state *GameState) ChangeState(indicator int) {
	var newLocation pixel.Matrix
	currentLocation := state.mainPlayer.GetObjectMatrix()
	switch indicator {
	case PlayerMovedLeft:
		newLocation = currentLocation.Moved(pixel.V(-utils.StepSize, 0))
	case PlayerMovedRight:
		newLocation = currentLocation.Moved(pixel.V(utils.StepSize, 0))
	case PlayerMovedDown:
		newLocation = currentLocation.Moved(pixel.V(0, -utils.StepSize))
	case PlayerMovedUp:
		newLocation = currentLocation.Moved(pixel.V(0, utils.StepSize))
	case PlayerShotBullet:
		// Start shooting...
	default:
		newLocation = currentLocation
	}
	if newLocation != currentLocation && inBounds(newLocation) {
		state.mainPlayer.SetMatrix(newLocation)
	}
}

func inBounds(mat pixel.Matrix) bool {
	if (mat[4] < utils.WindowWidth && mat[4] > 0) &&
		(mat[5] < utils.WindowHeight && mat[5] > 0) {
		return true
	}
	return false
}
