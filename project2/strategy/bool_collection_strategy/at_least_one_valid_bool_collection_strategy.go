package bool_collection_strategy

type atLeastOneValidBoolCollectionStrategy struct{}

func NewAtLeastOneValidStrategy() *atLeastOneValidBoolCollectionStrategy {
	return &atLeastOneValidBoolCollectionStrategy{}
}

func (alovs *atLeastOneValidBoolCollectionStrategy) CheckElements(booleans []bool) bool {

	for _, boolean := range booleans {
		if boolean {
			return true
		}
	}

	return false
}
