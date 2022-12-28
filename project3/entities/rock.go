package entities

type Rock struct {
	x int
	y int
}

func NewRock(x int, y int) *Rock {
	return &Rock{
		x: x,
		y: y,
	}
}

func (r *Rock) GetVisualString() string {
	return "O"
}

func (r *Rock) GetX() int {
	return r.x
}

func (r *Rock) GetY() int {
	return r.y
}

func (r *Rock) SetXY(x int, y int) {
	r.x = x
	r.y = y
}

func (t *Rock) IsPassable() bool {
	return false
}
