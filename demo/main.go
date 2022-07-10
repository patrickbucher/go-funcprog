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
	fmt.Println(g(10.0))
}
