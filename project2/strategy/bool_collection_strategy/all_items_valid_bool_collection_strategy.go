package bool_collection_strategy

type allItemsValidBoolCollectionStrategy struct{}

func NewAllItemsValidBoolCollectionStrategy() *allItemsValidBoolCollectionStrategy {
	return &allItemsValidBoolCollectionStrategy{}
}

func (aivs *allItemsValidBoolCollectionStrategy) CheckElements(booleans []bool) bool {

	for _, boolean := range booleans {
		if !boolean {
			return false
		}
	}

	return true
}
