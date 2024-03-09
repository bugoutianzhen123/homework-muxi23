package builder

type Builder[T any] struct {
	buffer []T
}

func (b *Builder[T]) Write(src []T) (n int, err error) {
	b.buffer = append(b.buffer, src...)
	return len(src), nil
}

func (b *Builder[T]) Read(dest []T) (n int, err error) {
	copy(dest, b.buffer)
	if len(dest) > len(b.buffer) {
		b.buffer = nil
		return len(b.buffer), nil
	}
	b.buffer = b.buffer[len(dest):]
	return len(dest), nil
}
