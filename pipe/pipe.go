package pipe

import (
	"rhin/lib"
)

func P1[A any, B any](
	value A,
	op1 func(input A) B,
	) B {
	return op1(value)
}

func P2[A any, B any, C any](
	value A,
	op1 func(input A) B,
	op2 func(input B) C,
	) C {
	return op2(op1(value))
}

func P3[A any, B any, C any, D any](
	value A,
	op1 func(input A) B,
	op2 func(input B) C,
	op3 func(input C) D,
	) D {
	return op3(op2(op1(value)))
}

func Filter[T any](condition func(item T) bool) func([]T) []T {
	return func(items []T) []T {
		return lib.Filter(items, condition)
	}
}

func Map[T any, X any](transform func(item T) X) func([]T) []X {
	return func(items []T) []X {
		return lib.Map(items, transform)
	}
}

func Every[T any](condition func(item T) bool) func([]T) bool {
	return func(items []T) bool {
		return lib.Every(items, condition)
	}
}

func Some[T any](condition func(item T) bool) func([]T) bool  {
	return func(items []T) bool {
		return lib.Some(items, condition)
	}
}