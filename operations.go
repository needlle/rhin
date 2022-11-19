package rhin

func Filter[T any](condition func(T) bool) func([]T) []T {
	return func(items []T) []T {
		result := []T{}
		for i := 0; i < len(items); i++ {
			if res := condition(items[i]); res {
				result = append(result, items[i])
			}
		}
		return result
	}
}

func Map[T any, X any](transform func(T) X) func([]T) []X {
	return func(items []T) []X {
		result := []X{}
		for i := 0; i < len(items); i++ {
			result = append(result, transform(items[i]))
		}
		return result
	}
}

func Reduce[T any, X any](reducer func(X, T) X, defaultValue X) func([]T) X {
	result := defaultValue
	return func(items []T) X {
		for i := 0; i < len(items); i++ {
			result = reducer(result, items[i])
		}
		return result
	}
}

func Every[T any](condition func(T) bool) func([]T) bool {
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
