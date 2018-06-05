package gamestate

import (
	"github.com/faiface/pixel"
	"github.com/nimrodshn/GoInvaders/gameobject"
	"github.com/nimrodshn/GoInvaders/utils"
	"time"
)

// GameState holds the current game state
type GameState struct {
	gameObjects          map[int][]*gameobject.GameObject
	lastTimeShot         time.Time
	timeBeginningOfLevel time.Time
	level                int
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
	levelInterval = time.Duration(10 * time.Second)
	// the amount of initial game enemies
	enemyCount = 5
)

// NewGameState Creates  a new GameState for game initialization
func NewGameState() (state *GameState, err error) {
	state = new(GameState)
	state.lastTimeShot = time.Now()
	state.timeBeginningOfLevel = time.Now()
	state.level = 1
	objects := make(map[int][]*gameobject.GameObject)

	// Initialize a new main player.
	mainPlayerLoctaion := pixel.V(float64(utils.WindowWidth/2), float64(utils.WindowHeight/10))
	player := gameobject.NewGameObject(mainPlayerLoctaion, gameobject.MainPlayerObject)
	objects[gameobject.MainPlayerObject] = []*gameobject.GameObject{player}
	// Initialize a new enemy player.
	enemies := initializeEnemiesForLevel(state.level)
	objects[gameobject.EnemyObject] = make([]*gameobject.GameObject, 0)
	objects[gameobject.EnemyObject] = append(objects[gameobject.EnemyObject], enemies...)

	state.gameObjects = objects
	return state, err
}

// ChangeState changes the game state according to the input given by ui.
func (state *GameState) ChangeState(indicator int) {
	state.processInput(indicator)
	state.gameObjects[gameobject.BulletObject] = updateObjectsLocation(state.gameObjects[gameobject.BulletObject], utils.StepSize)
	state.gameObjects[gameobject.EnemyObject] = updateObjectsLocation(state.gameObjects[gameobject.EnemyObject], -utils.StepSize/5)
	if time.Since(state.timeBeginningOfLevel) >= levelInterval {
		enemies := initializeEnemiesForLevel(1)
		state.gameObjects[gameobject.EnemyObject] = append(state.gameObjects[gameobject.EnemyObject], enemies...)
		state.timeBeginningOfLevel = time.Now()
	}
}

func (state *GameState) processInput(indicator int) {
	var newLocation pixel.Matrix
	mainPlayer := state.GetMainPlayer()
	playerMat := mainPlayer.GetObjectMatrix()
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
			b := gameobject.NewGameObject(playerVec, gameobject.BulletObject)
			bulletes := state.gameObjects[gameobject.BulletObject]
			bulletes = append(bulletes, b)
			state.gameObjects[gameobject.BulletObject] = bulletes
			state.lastTimeShot = time.Now()
		}
	}
	if newLocation != playerMat && inBounds(newLocation) {
		mainPlayer.SetMatrix(newLocation)
	}
}

func updateObjectsLocation(objects []*gameobject.GameObject, step float64) []*gameobject.GameObject {
	updatedObjects := make([]*gameobject.GameObject, 0)
	for _, b := range objects {
		mat := b.GetObjectMatrix()
		newLocation := mat.Moved(pixel.V(0, step))
		if inBounds(newLocation) {
			b.SetMatrix(newLocation)
			updatedObjects = append(updatedObjects, b)
		}
	}
	return updatedObjects
}

func inBounds(mat pixel.Matrix) bool {
	if (mat[4] < utils.WindowWidth && mat[4] > 0) &&
		(mat[5] < utils.WindowHeight && mat[5] > 0) {
		return true
	}
	return false
}

func initializeEnemiesForLevel(level int) []*gameobject.GameObject {
	total := level * enemyCount
	idx := 0
	enemyArr := make([]*gameobject.GameObject, total)
	for i := 0; i < 2*enemyCount; i++ {
		if i%2 != 0 {
			x := float64(i) / float64(2*enemyCount)
			location := pixel.V(x*utils.WindowWidth, utils.WindowHeight)
			e := gameobject.NewGameObject(location, gameobject.EnemyObject)
			enemyArr[idx] = e
			idx++
		}
	}
	return enemyArr
}

// GetMainPlayer returns the main player.
func (state *GameState) GetMainPlayer() *gameobject.GameObject {
	return state.gameObjects[gameobject.MainPlayerObject][0]
}

// GetEnemies returns the current enemies.
func (state *GameState) GetEnemies() []*gameobject.GameObject {
	return state.gameObjects[gameobject.EnemyObject]
}

// SetEnemies sets the current enemies.
func (state *GameState) SetEnemies(enemies []*gameobject.GameObject) {
	state.gameObjects[gameobject.EnemyObject] = enemies
}

// GetBullets returns the current bullets.
func (state *GameState) GetBullets() []*gameobject.GameObject {
	return state.gameObjects[gameobject.BulletObject]
}

// SetBullets sets the current bullets.
func (state *GameState) SetBullets(bullets []*gameobject.GameObject) {
	state.gameObjects[gameobject.BulletObject] = bullets
}
