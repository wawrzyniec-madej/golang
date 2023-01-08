package interfaces

type BoardIface interface {
	Print()
	Clear()
	GetSize() (int, int)
	Tick()
}
