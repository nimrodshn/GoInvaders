package gameobject

import (
	"github.com/faiface/pixel"
)

const (
	// BulletObject is constant representing the bullete object.
	BulletObject = 0
	// MainPlayerObject is a constant representing player object.
	MainPlayerObject = 1
	// EnemyObject is a constant representing enemy object.
	EnemyObject = 2
)

// GameObject abstracts away the game object.
type GameObject struct {
	mat        pixel.Matrix
	objectType int
}

// NewGameObject returns a new game object.
func NewGameObject(initLocation pixel.Vec, objectType int) *GameObject {
	mat := pixel.IM
	mat = mat.Moved(initLocation)
	object := new(GameObject)
	object.mat = mat
	object.objectType = objectType
	return object
}

// SetMatrix allows to set object matrix.
func (obj *GameObject) SetMatrix(matrix pixel.Matrix) {
	obj.mat = matrix
}

// GetObjectMatrix Returns the object matrix containing information needed
// in order to render.
func (obj *GameObject) GetObjectMatrix() pixel.Matrix {
	return obj.mat
}
