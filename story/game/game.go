package game

import (
	"story/blocks"
	"story/interfaces"
)

type game struct {
	InitialBlock interfaces.BlockInterface
}

func NewGame() *game {
	return &game{
		InitialBlock: blocks.NewTextBlock(
			"One night you decide to go on a stroll.",
			blocks.NewTextBlock(
				"While passing by the river, you see a tree stump.",
				blocks.NewChoiceBlock(
					"Will you sit on it?",
					[]interfaces.BlockInterface{
						blocks.NewTextBlock(
							"Yes!",
							blocks.NewTextBlock(
								"The stump is comfortable - you feel at one with the nature.",
								blocks.NewChoiceBlock(
									"What will you do now?",
									[]interfaces.BlockInterface{
										blocks.NewTextBlock(
											"Go back home.",
											blocks.NewTextBlock(
												"You headed back with the good memories of the stump.",
												nil,
											),
										),
										blocks.NewTextBlock(
											"Stay and enjoy the view.",
											blocks.NewTextBlock(
												"You fall asleep to the sound of the water.",
												blocks.NewTextBlock(
													"Next day you woke up in your bed with the good memories of the stump.",
													nil,
												),
											),
										),
									},
								),
							),
						),
						blocks.NewTextBlock(
							"No!",
							blocks.NewTextBlock(
								"You passed the stump and headed back home.",
								nil,
							),
						),
					},
				),
			),
		),
	}
}

func (g *game) Start() {
	g.InitialBlock.Progress()
}
