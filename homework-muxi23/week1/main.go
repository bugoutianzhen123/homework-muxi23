package main

import "fmt"

type Group[T any] interface {
	Write(t []T) (n int)
	Read(t []T) (n int)
}

type Builder[T any] struct {
	data []T
}

func (b *Builder[T]) Write(t []T) (n int) {
	b.data = append(b.data, t...)
	return len(t)
}

func (b *Builder[T]) Read(dest []T) (n int) {
	fmt.Println(b.data)
	copy(dest, b.data)
	if len(dest) > len(b.data) {
		b.data = nil
		return len(b.data)
	}
	b.data = b.data[len(dest):]
	return len(dest)
}

func main() {
	intbd := &Builder[int]{}
	var a Group[int]
	a = intbd
	t1 := a.Write([]int{1, 2, 3})
	fmt.Println(intbd.data)
	t2 := a.Read([]int{4, 56})
	fmt.Println(intbd.data)
	t3 := a.Write([]int{5, 5, 6})
	fmt.Println(intbd.data)
	t4 := a.Write([]int{7, 89})
	fmt.Println(intbd.data)
	fmt.Println(t1, t2, t3, t4)
}
