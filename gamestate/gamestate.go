package gamestate

import (
	"github.com/faiface/pixel"
	"github.com/nimrodshn/GoInvaders/bullete"
	"github.com/nimrodshn/GoInvaders/gameobject"
	"github.com/nimrodshn/GoInvaders/spaceship"
	"github.com/nimrodshn/GoInvaders/utils"
	"time"
)

// GameState holds the current game state
type GameState struct {
	mainPlayer       *spaceship.Spaceship
	enemies          []*spaceship.Spaceship
	bullets          []*bullete.Bullete
	lastTimeShot     time.Time
	timeForNextLevel time.Duration
	level            int
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
	// Interval from last shot, this is to prevent shoting storms.
	shotInterval = time.Duration(200 * time.Millisecond)
	// Time untill next deployment of enemies
	levelInterval = time.Duration(30 * time.Second)
	// the amount of initial game enemies
	enemyCount = 5
	// time between two levels
	timeForNextLevel = 1 * time.Minute
)

// NewGameState Creates  a new GameState for game initialization
func NewGameState() (state *GameState, err error) {
	state = new(GameState)
	player, err := spaceship.NewMainPlayer()
	state.mainPlayer = player
	state.lastTimeShot = time.Now()
	state.level = 1
	state.enemies = initializeEnemiesForLevel(state.level)
	return state, err
}

// GetGameObjects returns a snapshot of the current entities in the game.
func (state *GameState) GetGameObjects() []gameobject.GameObject {
	objects := make([]gameobject.GameObject, 0)
	// cannot append []T to and interface (see https://golang.org/doc/faq#convert_slice_of_interface).
	for _, enemy := range state.enemies {
		objects = append(objects, enemy)
	}
	for _, bullet := range state.bullets {
		objects = append(objects, bullet)
	}
	objects = append(objects, state.mainPlayer)
	return objects
}

// ChangeState changes the game state according to the input given by ui.
func (state *GameState) ChangeState(indicator int) {
	state.processInput(indicator)
	state.updateBulletesAndEnemiesLocation()
	state.ComputeLogic()
}

func (state *GameState) processInput(indicator int) {
	var newLocation pixel.Matrix
	playerMat := state.mainPlayer.GetObjectMatrix()
	switch indicator {
	case PlayerMovedLeft:
		newLocation = playerMat.Moved(pixel.V(-utils.StepSize, 0))
	case PlayerMovedRight:
		newLocation = playerMat.Moved(pixel.V(utils.StepSize, 0))
	case PlayerMovedDown:
		newLocation = playerMat.Moved(pixel.V(0, -utils.StepSize))
	case PlayerMovedUp:
		newLocation = playerMat.Moved(pixel.V(0, utils.StepSize))
	case PlayerShotBullet:
		if time.Since(state.lastTimeShot) >= shotInterval {
			playerVec := pixel.V(playerMat[4], playerMat[5])
			b, _ := bullete.NewBullete(playerVec)
			state.bullets = append(state.bullets, b)
			state.lastTimeShot = time.Now()
		}
	}
	if newLocation != playerMat && inBounds(newLocation) {
		state.mainPlayer.SetMatrix(newLocation)
	}
}

func (state *GameState) updateBulletesAndEnemiesLocation() {
	updatedBullets := make([]*bullete.Bullete, 0)
	updatedEnemies := make([]*spaceship.Spaceship, 0)
	for _, b := range state.bullets {
		bulleteMat := b.GetObjectMatrix()
		newLocation := bulleteMat.Moved(pixel.V(0, utils.StepSize))
		if inBounds(newLocation) {
			b.SetMatrix(newLocation)
			updatedBullets = append(updatedBullets, b)
		}
	}
	state.bullets = updatedBullets
	for _, e := range state.enemies {
		enemyeMat := e.GetObjectMatrix()
		newLocation := enemyeMat.Moved(pixel.V(0, -utils.StepSize/5))
		if inBounds(newLocation) {
			e.SetMatrix(newLocation)
			updatedEnemies = append(updatedEnemies, e)
		}
	}
	state.enemies = updatedEnemies
}

func inBounds(mat pixel.Matrix) bool {
	if (mat[4] < utils.WindowWidth && mat[4] > 0) &&
		(mat[5] < utils.WindowHeight && mat[5] > 0) {
		return true
	}
	return false
}

func initializeEnemiesForLevel(level int) []*spaceship.Spaceship {
	count := level * enemyCount
	idx := 0
	enemyArr := make([]*spaceship.Spaceship, count)
	for i := 0; i < 2*count; i++ {
		if i%2 != 0 {
			e, _ := spaceship.NewEnemy()
			mat := e.GetObjectMatrix()
			x := float64(i) / float64(2*count)
			e.SetMatrix(mat.Moved(pixel.V(x*utils.WindowWidth, utils.WindowHeight*0.8)))
			enemyArr[idx] = e
			idx++
		}
	}
	return enemyArr
}

// GetBulletes Returns game bullets
func (state *GameState) GetBulletes() []*bullete.Bullete {
	return state.bullets
}

// SetBulletes sets game bullets
func (state *GameState) SetBulletes(bulletes []*bullete.Bullete) {
	state.bullets = bulletes
}

// GetMainPlayer returns the main player.
func (state *GameState) GetMainPlayer() *spaceship.Spaceship {
	return state.mainPlayer
}

// GetEnemies returns the current enemies.
func (state *GameState) GetEnemies() []*spaceship.Spaceship {
	return state.enemies
}

// SetEnemies returns the current enemies.
func (state *GameState) SetEnemies(enemies []*spaceship.Spaceship) {
	state.enemies = enemies
}

// ComputeLogic computes the game logic as a lambda and changes the state in place.
func (state *GameState) ComputeLogic() {
	for j, b := range state.bullets {
		for i, e := range state.enemies {
			rect := e.GetObjectSprite().Frame()
			eMat := e.GetObjectMatrix()
			eX := eMat[4]
			eY := eMat[5]

			bMat := b.GetObjectMatrix()
			bX := bMat[4]
			bY := bMat[5]
			// If shot is contained within the borders of the enemy sprite.
			if bX >= (eX-rect.Max.X/2) && bX <= (eX+rect.Max.X/2) &&
				bY >= (eY-rect.Max.Y/2) && bY <= (eY+rect.Max.Y/2) {
				// Update bulletes
				state.bullets[j] = state.bullets[len(state.bullets)-1]
				state.bullets = state.bullets[:len(state.bullets)-1]

				// Update Enemies
				state.enemies[i] = state.enemies[len(state.enemies)-1]
				state.enemies = state.enemies[:len(state.enemies)-1]
			}
		}
	}
}
