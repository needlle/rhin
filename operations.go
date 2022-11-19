package rhin

func Filter[T any](condition func(item T) bool) func([]T) []T {
	return func(items []T) []T {
		newItems := []T{}
		for i := 0; i < len(items); i++ {
			if res := condition(items[i]); res {
				newItems = append(newItems, items[i])
			}
		}
		return newItems
	}
}

func Map[T any, X any](transform func(item T) X) func([]T) []X {
	return func(items []T) []X {
		newItems := []X{}
		for i := 0; i < len(items); i++ {
			newItems = append(newItems, transform(items[i]))
		}
		return newItems
	}
}

func Every[T any](condition func(item T) bool) func([]T) bool {
	return func(items []T) bool {
		for i := 0; i < len(items); i++ {
			if res := condition(items[i]); !res {
				return false
			}
		}
		return true
	}
}

func Some[T any](condition func(item T) bool) func([]T) bool {
	return func(items []T) bool {
		for i := 0; i < len(items); i++ {
			if res := condition(items[i]); res {
				return true
			}
		}
		return false
	}
}
