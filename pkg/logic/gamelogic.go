package logic

import (
	"github.com/faiface/pixel"
	"github.com/nimrodshn/GoInvaders/pkg/gamestate"
)

// ComputeLogic computes the game logic as a lambda and changes the state in place.
func ComputeLogic(state *gamestate.GameState, enemySprite *pixel.Sprite) {
	bullets := state.GetBullets()
	enemies := state.GetEnemies()
	for i, e := range enemies {
		eRect := enemySprite.Frame()
		eMat := e.GetObjectMatrix()
		for j, b := range bullets {
			bMat := b.GetObjectMatrix()

			if detectCollision(bMat, eMat, eRect) {
				bullets[j] = bullets[len(bullets)-1]
				state.SetBullets(bullets[:len(bullets)-1])

				// Update Enemies
				enemies[i] = enemies[len(enemies)-1]
				state.SetEnemies(enemies[:len(enemies)-1])
			}
		}
		mainPlayerMat := state.GetMainPlayer().GetObjectMatrix()
		if detectCollision(mainPlayerMat, eMat, eRect) {
			// Update Enemies
			enemies[i] = enemies[len(enemies)-1]
			state.SetEnemies(enemies[:len(enemies)-1])

			state.DecrementMainPlayerLives()
		}
	}
}

// TODO: make this collision detection better - we need to account for better collision not just when the
// center of item (represented by itemMatrix) lies inside the sprite represented by enemyRect.
// (This is the case now with main player!)
func detectCollision(itemMatrix pixel.Matrix, enemyMatrix pixel.Matrix, enemyRect pixel.Rect) bool {
	eX := enemyMatrix[4]
	eY := enemyMatrix[5]
	bX := itemMatrix[4]
	bY := itemMatrix[5]

	if bX >= (eX-enemyRect.Max.X/2) && bX <= (eX+enemyRect.Max.X/2) &&
		bY >= (eY-enemyRect.Max.Y/2) && bY <= (eY+enemyRect.Max.Y/2) {
		return true
	}

	return false
}
