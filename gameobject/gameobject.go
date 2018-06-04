package gameobject

import (
	"github.com/faiface/pixel"
)

// GameObject is an interface for game objects
// ment for ease of rendering
type GameObject interface {
	GetObjectMatrix() pixel.Matrix
	GetObjectSprite() *pixel.Sprite
	SetMatrix(mat pixel.Matrix)
}
