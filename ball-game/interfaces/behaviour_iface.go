package interfaces

type BehaviourIface interface {
	React(actionName string, entity BoardEntityIface, board BoardIface)
}
