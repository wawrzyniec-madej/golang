package blocks

import (
	"fmt"
	"story/interfaces"
)

type choiceBlock struct {
	Message    string
	NextBlocks []interfaces.BlockInterface
}

func NewChoiceBlock(message string, next_blocks []interfaces.BlockInterface) *choiceBlock {
	return &choiceBlock{
		Message:    message,
		NextBlocks: next_blocks,
	}
}

func (cb *choiceBlock) Progress() {
	fmt.Println(cb.Message)

	if len(cb.NextBlocks) == 0 {
		return
	}

	for idx, next_block := range cb.NextBlocks {
		fmt.Printf("%d: %s\n", idx, next_block.GetMessage())
	}

	cb.promptNextBlock()
}

func (cb *choiceBlock) promptNextBlock() {

	var selectedBlock interfaces.BlockInterface
	for {
		fmt.Print("Choose one option: ")

		var selectedIndex int
		fmt.Scanln(&selectedIndex)

		fmt.Println()

		if len(cb.NextBlocks)-1 < selectedIndex {
			fmt.Printf("Option '%d' is not valid.\n", selectedIndex)
			continue
		}

		block := cb.NextBlocks[selectedIndex]
		if block != nil {
			selectedBlock = block
			break
		}
	}

	selectedBlock.Progress()
}

func (cb *choiceBlock) GetMessage() string {
	return cb.Message
}
