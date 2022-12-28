package turtle_game

import (
	"fmt"
	"project3/entities"
	"project3/interfaces"

	"github.com/eiannone/keyboard"
)

const BOARD_SIZE int = 10

type TurtleGame struct {
	entities []interfaces.EntityInterface
}

func (tg *TurtleGame) Start() {

	pointerToTurtle := tg.spawnTurtle()
	tg.spawnRock(2, 2)
	tg.spawnRock(2, 3)

	for {
		tg.printBoard()
		tg.promptNextKey(pointerToTurtle)
	}
}

func (tg *TurtleGame) spawnTurtle() *entities.Turtle {

	turtle := entities.NewTurtle(0, 0)

	tg.entities = append(tg.entities, turtle)

	return turtle
}

func (tg *TurtleGame) spawnRock(x int, y int) {

	rock := entities.NewRock(x, y)

	tg.entities = append(tg.entities, rock)
}

func (tg *TurtleGame) promptNextKey(pointerToTurtle *entities.Turtle) {

	_, key, err := keyboard.GetSingleKey()
	if err != nil {
		panic(err)
	}

	if key == keyboard.KeyEsc {
		panic("game is done")
	}

	tg.handleTurtleMovement(
		pointerToTurtle,
		&key,
	)
}

func (tg *TurtleGame) handleTurtleMovement(pointerToTurtle *entities.Turtle, pointerToKey *keyboard.Key) {

	key := *pointerToKey

	var newX, newY int

	switch key {
	case keyboard.KeyArrowRight:
		newX = pointerToTurtle.GetX() + 1
		newY = pointerToTurtle.GetY()

	case keyboard.KeyArrowLeft:
		newX = pointerToTurtle.GetX() - 1
		newY = pointerToTurtle.GetY()

	case keyboard.KeyArrowUp:
		newX = pointerToTurtle.GetX()
		newY = pointerToTurtle.GetY() - 1

	case keyboard.KeyArrowDown:
		newX = pointerToTurtle.GetX()
		newY = pointerToTurtle.GetY() + 1
	}

	if tg.areValidXY(newX, newY) {
		entityAtXY := tg.getEntityAtCoords(newX, newY)

		if entityAtXY != nil && !entityAtXY.IsPassable() {
			return
		}

		fmt.Printf("Turtle moving to: %d, %d\n", newX, newY)
		pointerToTurtle.SetXY(newX, newY)
	}
}

func (tg *TurtleGame) printBoard() {
	for y := 0; y < BOARD_SIZE; y++ {
		for x := 0; x < BOARD_SIZE; x++ {

			entity_at_coords := tg.getEntityAtCoords(x, y)

			if entity_at_coords == nil {
				fmt.Print("·" + " ")
				continue
			}

			fmt.Print(entity_at_coords.GetVisualString() + " ")
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func (tg *TurtleGame) getEntityAtCoords(x int, y int) interfaces.EntityInterface {
	for _, entity := range tg.entities {
		if entity.GetX() == x && entity.GetY() == y {
			return entity
		}
	}

	return nil
}

func (tg *TurtleGame) areValidXY(newX int, newY int) bool {
	if newX < 0 || newX > BOARD_SIZE-1 {
		return false
	}

	if newY < 0 || newY > BOARD_SIZE-1 {
		return false
	}

	return true
}
