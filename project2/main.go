package main

import (
	"fmt"
	"project2/chains/validator_chain"
	"project2/interfaces"
	"project2/prompters"
	"project2/strategy/bool_collection_strategy"
	"project2/validators/text_validator"
)

func main() {

	for {

		text := prompters.Prompter{}.Prompt("Podaj pierwszą wartość")

		validators := []interfaces.TextValidatorInterface{
			text_validator.NewHasExclamationTextValidator(),
			text_validator.NewIsAbbaTextValidator(),
		}

		validatorChain := validator_chain.NewTextValidatorChain(
			validators,
			bool_collection_strategy.NewAtLeastOneValidStrategy(),
		)

		if !validatorChain.Validate(text) {
			fmt.Printf("Wartość %s spełnia warunki walidacji.\n", text)
			continue
		}

		fmt.Printf("Wartość %s nie spełnia warunków waliacji.\n", text)
		break
	}
}
