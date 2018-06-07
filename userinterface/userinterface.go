package userinterface

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/nimrodshn/GoInvaders/bullete"
	"github.com/nimrodshn/GoInvaders/gamestate"
	"github.com/nimrodshn/GoInvaders/logic"
	"github.com/nimrodshn/GoInvaders/spaceship"
	"github.com/nimrodshn/GoInvaders/utils"
	"golang.org/x/image/colornames"
)

// UserInterface represents the Interface which takes input from user
// and displays output accordingly.
// UserInterface also updates the game state according to the input.
type UserInterface struct {
	state            *gamestate.GameState
	window           *pixelgl.Window
	enemySprite      *spaceship.Spaceship
	bulletSprite     *bullete.Bullete
	mainPlayerSprite *spaceship.Spaceship
}

// NewUserInterface creates a user interface object
func NewUserInterface() (UserInterface, error) {
	cfg := pixelgl.WindowConfig{
		Title:  "GoInvaders",
		Bounds: pixel.R(0, 0, utils.WindowWidth, utils.WindowHeight),
		VSync:  true,
	}

	win, _ := pixelgl.NewWindow(cfg)
	state, _ := gamestate.NewGameState()
	player, _ := spaceship.NewMainPlayer()
	enemy, _ := spaceship.NewEnemy()
	bullet, _ := bullete.NewBullete()

	ui := UserInterface{
		window:           win,
		state:            state,
		mainPlayerSprite: player,
		enemySprite:      enemy,
		bulletSprite:     bullet,
	}
	return ui, nil
}

// DrawGameOnScreen draws the game on screen
func (ui *UserInterface) DrawGameOnScreen() {
	ui.window.Clear(colornames.Black)
	mainPlayer := ui.state.GetMainPlayer()
	enemies := ui.state.GetEnemies()
	bulletes := ui.state.GetBullets()

	bulletBatch := pixel.NewBatch(&pixel.TrianglesData{}, ui.bulletSprite.GetObjectSprite().Picture())
	mainPlayerBatch := pixel.NewBatch(&pixel.TrianglesData{}, ui.mainPlayerSprite.GetObjectSprite().Picture())
	enemiesBatch := pixel.NewBatch(&pixel.TrianglesData{}, ui.enemySprite.GetObjectSprite().Picture())

	// Draw bullets
	bulletBatch.Clear()
	for _, b := range bulletes {
		ui.bulletSprite.GetObjectSprite().Draw(bulletBatch, b.GetObjectMatrix())
	}
	bulletBatch.Draw(ui.window)

	// Draw main player
	mainPlayerBatch.Clear()
	ui.mainPlayerSprite.GetObjectSprite().Draw(mainPlayerBatch, mainPlayer.GetObjectMatrix())
	mainPlayerBatch.Draw(ui.window)

	// Draw enemies
	enemiesBatch.Clear()
	for _, e := range enemies {
		ui.enemySprite.GetObjectSprite().Draw(enemiesBatch, e.GetObjectMatrix())
	}
	enemiesBatch.Draw(ui.window)

	ui.window.Update()
}

// ListenOnKeyStroke Moves player on key strokes
func (ui UserInterface) ListenOnKeyStroke() {
	var userInput int
	switch {
	case ui.window.Pressed(pixelgl.KeyLeft):
		userInput = gamestate.PlayerMovedLeft
	case ui.window.Pressed(pixelgl.KeyRight):
		userInput = gamestate.PlayerMovedRight
	case ui.window.Pressed(pixelgl.KeyDown):
		userInput = gamestate.PlayerMovedDown
	case ui.window.Pressed(pixelgl.KeyUp):
		userInput = gamestate.PlayerMovedUp
	case ui.window.Pressed(pixelgl.KeySpace):
		userInput = gamestate.PlayerShotBullet
	}
	ui.state.ChangeState(userInput)
	logic.ComputeLogic(ui.state, ui.enemySprite.GetObjectSprite())
}

// WindowClosed Check if window is closed
func (ui UserInterface) WindowClosed() bool {
	if ui.window.Closed() {
		return true
	}
	return false
}
