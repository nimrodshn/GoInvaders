package logic

import (
	"github.com/faiface/pixel"
	"github.com/nimrodshn/GoInvaders/gamestate"
)

// ComputeLogic computes the game logic as a lambda and changes the state in place.
func ComputeLogic(state *gamestate.GameState, enemySprite *pixel.Sprite) {
	bullets := state.GetBullets()
	enemies := state.GetEnemies()
	for j, b := range bullets {
		for i, e := range enemies {
			rect := enemySprite.Frame()
			eMat := e.GetObjectMatrix()

			bMat := b.GetObjectMatrix()

			// If shot is contained within the borders of the enemy sprite.
			if didBulletHit(bMat, eMat, rect) {
				bullets[j] = bullets[len(bullets)-1]
				state.SetBullets(bullets[:len(bullets)-1])

				// Update Enemies
				enemies[i] = enemies[len(enemies)-1]
				state.SetEnemies(enemies[:len(enemies)-1])
			}
		}
	}
}

func didBulletHit(bulletMatrix pixel.Matrix, enemyMatrix pixel.Matrix, enemyRect pixel.Rect) bool {
	eX := enemyMatrix[4]
	eY := enemyMatrix[5]
	bX := bulletMatrix[4]
	bY := bulletMatrix[5]

	if bX >= (eX-enemyRect.Max.X/2) && bX <= (eX+enemyRect.Max.X/2) &&
		bY >= (eY-enemyRect.Max.Y/2) && bY <= (eY+enemyRect.Max.Y/2) {
		return true
	}
	return false
}
