package lib

func Every[T any](items []T, condition func(item T) bool) bool {
	for i := 0; i < len(items); i++ {
		if res := condition(items[i]); !res {
			return false
		}
	}
	return true
}

func Some[T any](items []T, condition func(item T) bool) bool {
	for i := 0; i < len(items); i++ {
		if res := condition(items[i]); res {
			return true
		}
	}
	return false
}

func Filter[T any](items []T, condition func(item T) bool) []T {
	newItems := []T{}
	for i := 0; i < len(items); i++ {
		if res := condition(items[i]); res {
			newItems = append(newItems, items[i])
		}
	}
	return newItems
}

func Map[T any, X any](items []T, transform func(item T) X) []X {
	newItems := []X{}
	for i := 0; i < len(items); i++ {
		newItems = append(newItems, transform(items[i]))
	}
	return newItems
}
