package blocks

import (
	"fmt"
	"story/interfaces"
)

type textBlock struct {
	Message   string
	NextBlock interfaces.BlockInterface
}

func NewTextBlock(message string, next_block interfaces.BlockInterface) *textBlock {
	return &textBlock{
		Message:   message,
		NextBlock: next_block,
	}
}

func (tb *textBlock) Progress() {
	fmt.Println(tb.Message)
	fmt.Scanln()

	if tb.NextBlock == nil {
		return
	}

	tb.NextBlock.Progress()
}

func (tb *textBlock) GetMessage() string {
	return tb.Message
}
