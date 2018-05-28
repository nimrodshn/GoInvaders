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

// ListenAndMoveOnKeyStroke Moves player on key strokes
func (ui UserInterface) ListenOnKeyStroke() {
	var newLocation pixel.Matrix
	currentLocation := ui.state.GetMainPlayer().GetObjectMatrix()
	switch {
	case ui.window.Pressed(pixelgl.KeyLeft):
		newLocation = currentLocation.Moved(pixel.V(-utils.StepSize, 0))
	case ui.window.Pressed(pixelgl.KeyRight):
		newLocation = currentLocation.Moved(pixel.V(utils.StepSize, 0))
	case ui.window.Pressed(pixelgl.KeyDown):
		newLocation = currentLocation.Moved(pixel.V(0, -utils.StepSize))
	case ui.window.Pressed(pixelgl.KeyUp):
		newLocation = currentLocation.Moved(pixel.V(0, utils.StepSize))
	case ui.window.Pressed(pixelgl.KeySpace):
		newLocation = currentLocation
	}
	if currentLocation != newLocation && inBounds(newLocation) {
		ui.state.ChangePlayerState(newLocation)
	}
}

func inBounds(mat pixel.Matrix) bool {
	if (mat[4] < utils.WindowWidth && mat[4] > 0) &&
		(mat[5] < utils.WindowHeight && mat[5] > 0) {
		return true
	}
	return false
}

func (ui UserInterface) WindowClosed() bool {
	if ui.window.Closed() {
		return true
	}
	return false
}
