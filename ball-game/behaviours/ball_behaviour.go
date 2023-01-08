package behaviours

import (
	"ball-game/enums"
	"ball-game/interfaces"
)

type BallBehaviour struct{}

func (bb *BallBehaviour) React(actionName string, entity interfaces.BoardEntityIface, board interfaces.BoardIface) {
	ball, ok := entity.(interfaces.BallIface)

	if !ok {
		return
	}

	if actionName == enums.ACTION_NAME_TICK {
		bb.onTick(ball, board)
	}
}

func (bb *BallBehaviour) onTick(ball interfaces.BallIface, board interfaces.BoardIface) {

	boardSizeX, boardSizeY := board.GetSize()

	entityPosX, entityPosY := ball.GetPos()

	moveVectorX, moveVectorY := ball.GetMoveVector()

	nextEntityPosX := entityPosX + moveVectorX
	nextEntityPosY := entityPosY + moveVectorY

	if nextEntityPosX > boardSizeX || nextEntityPosX < 0 {
		moveVectorX = moveVectorX * -1
	}

	if nextEntityPosY > boardSizeY || nextEntityPosY < 0 {
		moveVectorY = moveVectorY * -1
	}

	ball.SetPos(
		entityPosX+moveVectorX,
		entityPosY+moveVectorY,
	)

	ball.SetMoveVector(moveVectorX, moveVectorY)
}
