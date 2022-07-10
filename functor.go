package funcprog

type ListFunctor[T any] []T

func (l ListFunctor[T]) Map(f ErrFunc[T]) ListFunctor[T] {
	var ys ListFunctor[T]
	for _, x := range l {
		y, _ := f(x)
		ys = append(ys, y)
	}
	return ys
}
