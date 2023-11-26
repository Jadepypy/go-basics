package slice

type ISlice[T comparable] interface {
	Append(val T)
	// Delete removes the elements slice[i:j]
	Delete(idxFrom int, idxTo int)
	Get(idx int) T
}

type Slice[T comparable] struct {
	data []T
}

func NewSlice[T comparable](vals ...T) *Slice[T] {
	s := Slice[T]{}
	s.data = make([]T, len(vals))
	for i, v := range vals {
		s.data[i] = v
	}
	return &s
}

func (s *Slice[T]) Get(idx int) T {
	if idx >= s.Len() {
		panic("index exceeds the length of slice")
	}
	return s.data[idx]
}

func (s *Slice[T]) Len() int {
	return len(s.data)
}

func (s *Slice[T]) Append(val T) {
	s.data = append(s.data, val)
}

func (s *Slice[T]) Delete(idxFrom int, idxTo int) {
	_ = s.data[idxFrom:idxTo]

	m := len(s.data)
	n := m - (idxTo - idxFrom + 1)
	if 2*n < cap(s.data) {
		s2 := make([]T, n, cap(s.data)/2)
		copy(s2, s.data[:idxFrom])
		copy(s2[idxFrom:], s.data[idxTo+1:])
		s.data = s2
		return
	}
	s.data = append(s.data[:idxFrom], s.data[idxTo+1:]...)
	//fmt.Println("cap after, should no change", cap(s.data))

}

func (s *Slice[T]) Cap() int {
	return cap(s.data)
}
