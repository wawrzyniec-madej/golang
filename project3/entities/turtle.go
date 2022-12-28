package entities

type Turtle struct {
	x int
	y int
}

func NewTurtle(x int, y int) *Turtle {
	return &Turtle{
		x: x,
		y: y,
	}
}

func (t *Turtle) GetVisualString() string {
	return "T"
}

func (t *Turtle) GetX() int {
	return t.x
}

func (t *Turtle) GetY() int {
	return t.y
}

func (t *Turtle) SetXY(x int, y int) {
	t.x = x
	t.y = y
}

func (t *Turtle) IsPassable() bool {
	return true
}
