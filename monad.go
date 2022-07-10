package funcprog

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type ListMonad[T any] []T

type LiftingFunction[T any] func(T) ListMonad[T]

func Lift[T any](x T) ListMonad[T] {
	return []T{x}
}

func (lm ListMonad[T]) Flatten() T {
	return lm[0]
}

func (lm ListMonad[T]) FlatMap(f LiftingFunction[T]) ListMonad[T] {
	ys := make([]T, len(lm))
	for i, x := range lm {
		ys[i] = f(x).Flatten()
	}
	return ys
}

func ComposeLifting[T any](fs ...LiftingFunction[T]) LiftingFunction[T] {
	return func(x T) ListMonad[T] {
		var y ListMonad[T]
		for i := len(fs) - 1; i >= 0; i-- {
			f := fs[i]
			y = f(x)
			x = y.Flatten()
		}
		return y
	}
}

func Increment[T constraints.Integer](x T) ListMonad[T] {
	return Lift(x + 1)
}

func Twice[T constraints.Integer](x T) ListMonad[T] {
	return Lift(x * 2)
}

func Demo() {
	increment := Increment[int]
	twice := Twice[int]
	fmt.Println(Lift(3).FlatMap(increment).FlatMap(twice))

	numbers := ListMonad[int]{1, 2, 3}
	fmt.Println(numbers.FlatMap(increment).FlatMap(twice))

	f := ComposeLifting(increment, twice)
	fmt.Println(numbers.FlatMap(f))
}
