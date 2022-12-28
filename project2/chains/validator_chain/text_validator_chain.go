package validator_chain

import "project2/interfaces"

type textValidatorChain struct {
	validators []interfaces.TextValidatorInterface
	strategy   interfaces.BoolCollectionStrategyInterface
}

func NewTextValidatorChain(validators []interfaces.TextValidatorInterface, strategy interfaces.BoolCollectionStrategyInterface) *textValidatorChain {
	return &textValidatorChain{
		validators: validators,
		strategy:   strategy,
	}
}

func (tvc *textValidatorChain) Validate(text string) bool {

	results := []bool{}

	for _, validator := range tvc.validators {

		results = append(results, validator.IsValid(text))
	}

	return tvc.strategy.CheckElements(results)
}
