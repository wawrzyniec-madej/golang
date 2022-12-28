package types

type coords struct {
	x int
	y int
}

func New(x int, y int) *coords {
	return &coords{
		x: x,
		y: y,
	}
}
