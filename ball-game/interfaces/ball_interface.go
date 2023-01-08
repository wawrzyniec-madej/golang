package interfaces

type BallIface interface {
	BoardEntityIface
	GetMoveVector() (int, int)
	SetMoveVector(moveVectorX int, moveVectorY int)
}
