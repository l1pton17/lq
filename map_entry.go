package lq

type MapEntry[TKey, TValue any] struct {
	Key   TKey
	Value TValue
}

func NewMapEntry[TKey, TValue any](k TKey, v TValue) MapEntry[TKey, TValue] {
	return MapEntry[TKey, TValue]{
		Key:   k,
		Value: v,
	}
}
