package userinterface

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/nimrodshn/GoInvaders/gameobject"
	"github.com/nimrodshn/GoInvaders/gamestate"
	"github.com/nimrodshn/GoInvaders/utils"
	"golang.org/x/image/colornames"
)

// UserInterface represents the Interface which takes input from user
// and displays output accordingly.
// UserInterface also updates the game state according to the input.
type UserInterface struct {
	state  *gamestate.GameState
	window *pixelgl.Window
}

func NewUserInterface() (UserInterface, error) {
	cfg := pixelgl.WindowConfig{
		Title:  "GoInvaders",
		Bounds: pixel.R(0, 0, utils.WindowWidth, utils.WindowHeight),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	state, err := gamestate.NewGameState()
	if err != nil {
		panic(err)
	}
	ui := UserInterface{
		window: win,
		state:  state,
	}
	return ui, err
}

// DrawGameOnScreen draws the game on screen
func (ui *UserInterface) DrawGameOnScreen() {
	ui.window.Clear(colornames.Black)
	gameObjects := ui.state.GetGameObjects()
	for _, obj := range gameObjects {
		drawObjectOnScreen(obj, ui.window)
	}
	ui.window.Update()
}

func drawObjectOnScreen(object gameobject.GameObject, window *pixelgl.Window) {
	sprite := object.GetObjectSprite()
	mat := object.GetObjectMatrix()
	sprite.Draw(window, mat)
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
}

// WindowClosed Check if window is closed
func (ui UserInterface) WindowClosed() bool {
	if ui.window.Closed() {
		return true
	}
	return false
}
