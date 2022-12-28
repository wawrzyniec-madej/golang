package interfaces

type EntityInterface interface {
	GetVisualString() string
	GetX() int
	GetY() int
	SetXY(x int, y int)
	IsPassable() bool
}
