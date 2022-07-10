package funcprog

type Functor[T any] interface {
	Map(ErrFunc[T]) Functor[T]
}

type List[T any] []T

func (l List[T]) Map(f ErrFunc[T]) Functor[T] {
	var ys List[T]
	for _, x := range l {
		y, _ := f(x)
		ys = append(ys, y)
	}
	return ys
}
