package funcprog

import (
	"errors"

	"golang.org/x/exp/constraints"
)

type ErrFunc[T any] func(T) (T, error)

func Compose[T any](fs ...ErrFunc[T]) ErrFunc[T] {
	return func(x T) (T, error) {
		var y T
		var err error
		for i := len(fs) - 1; i >= 0; i-- {
			f := fs[i]
			y, err = f(x)
			if err != nil {
				return *new(T), err
			}
			x = y
		}
		return y, nil
	}
}

func Add[T constraints.Float](y T) ErrFunc[T] {
	return func(x T) (T, error) {
		return x + y, nil
	}
}

func Sub[T constraints.Float](y T) ErrFunc[T] {
	return func(x T) (T, error) {
		return x - y, nil
	}
}

func Mul[T constraints.Float](y T) ErrFunc[T] {
	return func(x T) (T, error) {
		return x * y, nil
	}
}

var ErrDivideByZero = errors.New("divide by zero")

func Div[T constraints.Float](y T) ErrFunc[T] {
	return func(x T) (T, error) {
		if y == 0 {
			return *new(T), ErrDivideByZero
		}
		return x / y, nil
	}
}
