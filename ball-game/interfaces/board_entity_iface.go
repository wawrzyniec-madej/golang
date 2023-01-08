package interfaces

type BoardEntityIface interface {
	GetPos() (int, int)
	GetRepresentation() string
	React(actionName string, board BoardIface)
	SetPos(posX int, poxY int)
}
