package main

import (
	"fmt"

	fp "github.com/patrickbucher/funcprog"
)

func main() {
	increment := fp.Add(1.0)
	decrement := fp.Sub(1.0)
	twice := fp.Mul(2.0)
	half := fp.Div(2.0)

	f := fp.Compose(increment, twice, decrement, half)
	fmt.Println(f(10.0)) // (((10 + 1) * 2) - 1) / 2 = 10.5

	byZero := fp.Div(0.0)
	g := fp.Compose(byZero, f)
	fmt.Println(g(10.0)) // divide by zero

	out := func(x float64) (float64, error) {
		fmt.Println(x)
		return x, nil
	}
	h := fp.Compose(increment, out, twice, out, decrement, out, half, out)
	fmt.Println(h(10.0)) // same as f, byt with logging

	numbers := fp.ListFunctor[float64]{1.0, 2.0, 3.0}

	// First Functor Law: Identity Function
	fmt.Println(numbers.Map(func(x float64) (float64, error) { return x, nil }))

	// Second Functor Law: f as a composition of other functions
	fmt.Println(numbers.Map(increment).Map(twice).Map(decrement).Map(half))
	fmt.Println(numbers.Map(f))

	fp.Demo()
}
