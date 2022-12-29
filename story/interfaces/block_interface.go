package interfaces

type BlockInterface interface {
	Progress()
	GetMessage() string
}
