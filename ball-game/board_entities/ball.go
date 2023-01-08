package board_entities

import (
	"ball-game/behaviours"
	"ball-game/interfaces"
)

type ball struct {
	posX           int
	posY           int
	representation string
	behaviours     []interfaces.BehaviourIface
	moveVectorX    int
	moveVectorY    int
}

func NewBall(posX int, posY int) *ball {
	return &ball{
		posX:           posX,
		posY:           posY,
		representation: "●",
		behaviours: []interfaces.BehaviourIface{
			&behaviours.BallBehaviour{},
		},
		moveVectorX: 1,
		moveVectorY: 1,
	}
}

func (b *ball) GetPos() (int, int) {
	return b.posX, b.posY
}

func (b *ball) SetPos(posX int, posY int) {
	b.posX = posX
	b.posY = posY
}

func (b *ball) GetRepresentation() string {
	return b.representation
}

func (b *ball) React(actionName string, board interfaces.BoardIface) {
	for _, behaviour := range b.behaviours {
		behaviour.React(actionName, b, board)
	}
}

func (b *ball) GetMoveVector() (int, int) {
	return b.moveVectorX, b.moveVectorY
}

func (b *ball) SetMoveVector(moveVectorX int, moveVectorY int) {
	b.moveVectorX = moveVectorX
	b.moveVectorY = moveVectorY
}
