package main

import (
	"fmt"
	"github.com/big-dust/homework-muxi23/week01/builder"
)

type Reader[T any] interface {
	Read(dest []T) (n int, err error)
}

type Writer[T any] interface {
	Write(src []T) (n int, err error)
}

type ReaderWriter[T any] interface {
	Reader[T]
	Writer[T]
}

func main() {
	b := &builder.Builder[byte]{}
	var r Reader[byte]
	r = b
	w := r.(Writer[byte])
	rw := r.(ReaderWriter[byte])
	n1, _ := rw.Write([]byte{1, 2, 3})
	n2, _ := w.Write([]byte{4, 5, 6})
	dest := make([]byte, n1+n2)
	n3, _ := r.Read(dest)
	fmt.Println(n1, n2, n3, dest)
}
