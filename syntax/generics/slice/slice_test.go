package slice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Slice_Delete(t *testing.T) {
	testCases := []struct {
		name          string
		s             *Slice[int]
		expectedSlice *Slice[int]
		expectedCap   int
		deletedIdx    []int
	}{
		{
			name:          "delete the first element",
			s:             NewSlice[int](1, 2, 3),
			expectedSlice: NewSlice[int](2, 3),
			expectedCap:   3,
			deletedIdx:    []int{0, 0},
		},
		{
			name:          "delete the last element",
			s:             NewSlice[int](1, 2, 3),
			expectedSlice: NewSlice[int](1, 2),
			expectedCap:   3,
			deletedIdx:    []int{2, 2},
		},
		{
			name:          "delete the middle element",
			s:             NewSlice[int](1, 2, 3),
			expectedSlice: NewSlice[int](1, 3),
			expectedCap:   3,
			deletedIdx:    []int{1, 1},
		},
		{
			name:          "delete the last two element",
			s:             NewSlice[int](1, 2, 3),
			expectedSlice: NewSlice[int](1),
			expectedCap:   1,
			deletedIdx:    []int{1, 2},
		},
		{
			name:          "reduce the capacity by half",
			s:             NewSlice[int](1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8),
			expectedSlice: NewSlice[int](1, 2),
			expectedCap:   16,
			deletedIdx:    []int{2, 31},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.s.Delete(tc.deletedIdx[0], tc.deletedIdx[1])
			assert.Equal(t, tc.expectedSlice, tc.s)
			assert.Equal(t, tc.expectedCap, tc.s.Cap())
		})
	}
}

func Test_Slice_Append(t *testing.T) {
	a := NewSlice[int]()
	a.Append(1)
	b := NewSlice[int](1, 2, 3)
	isEqual, msg := Equal(a, b)
	assert.Equal(t, isEqual, false)
	assert.Equal(t, msg, "the length of slice is different")

	a = NewSlice[int]()
	a.Append(1)
	a.Append(3)
	a.Append(3)
	fmt.Printf("%+v\n", a)
	b = NewSlice[int](1, 2, 3)
	isEqual, msg = Equal(a, b)
	assert.Equal(t, isEqual, false)
	assert.Equal(t, msg, "index: 1, value: 3 and value: 2 does not match")

	a = NewSlice[int](1, 2, 3)
	b = NewSlice[int](1, 2, 3)
	isEqual, msg = Equal(a, b)
	assert.Equal(t, isEqual, true)
}

func Equal[T comparable](a *Slice[T], b *Slice[T]) (bool, string) {
	if a.Len() != b.Len() {
		return false, "the length of slice is different"
	}

	for i := 0; i < a.Len(); i++ {
		if a.data[i] != b.data[i] {
			return false, fmt.Sprintf("index: %d, value: %+v and value: %+v does not match", i, a.Get(i), b.Get(i))
		}
	}
	return true, ""
}
