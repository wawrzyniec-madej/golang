package board

import (
	"ball-game/board_entities"
	"ball-game/enums"
	"ball-game/interfaces"
	"fmt"
	"os"
	"os/exec"
)

type Board struct {
	sizeX    int
	sizeY    int
	entities []interfaces.BoardEntityIface
}

func NewBoard(sizeX int, sizeY int) *Board {
	return &Board{
		sizeX: sizeX - 1,
		sizeY: sizeY - 1,
		entities: []interfaces.BoardEntityIface{
			board_entities.NewBall(0, 0),
			board_entities.NewBall(3, 0),
		},
	}
}

func (b *Board) GetSize() (int, int) {
	return b.sizeX, b.sizeY
}

func (b *Board) Print() {

	for y := 0; y <= b.sizeY; y++ {
		for x := 0; x <= b.sizeX; x++ {

			entityAtPos := b.getEntityAtPos(x, y)
			if entityAtPos != nil {
				fmt.Printf(" %s", entityAtPos.GetRepresentation())
			} else {
				fmt.Print(" □")
			}

			if x == b.sizeX {
				fmt.Print("\n")
			}
		}
	}
}

func (b *Board) Clear() {
	cmd := exec.Command("clear")

	cmd.Stdout = os.Stdout

	err := cmd.Run()

	if err != nil {
		panic("ERROR")
	}
}

func (b *Board) getEntityAtPos(posX int, posY int) interfaces.BoardEntityIface {
	for _, entity := range b.entities {
		entityPosX, entityPosY := entity.GetPos()

		if entityPosX == posX && entityPosY == posY {
			return entity
		}
	}

	return nil
}

func (b *Board) passActionName(actionName string) {

	for _, entity := range b.entities {
		entity.React(actionName, b)
	}
}

func (b *Board) Tick() {

	b.Clear()
	b.Print()
	b.passActionName(enums.ACTION_NAME_TICK)
}
