package ranges

type sliceRange[T any] []T

func (r *sliceRange[T]) Empty() bool           { return r == nil || len(*r) == 0 }
func (r *sliceRange[T]) Front() T              { return (*r)[0] }
func (r *sliceRange[T]) PopFront()             { *r = (*r)[1:] }
func (r *sliceRange[T]) Save() ForwardRange[T] { return SliceRange([]T(*r)) }

// SliceRange creates a range from a slice.
func SliceRange[T any](slice []T) ForwardRange[T] {
	result := sliceRange[T](slice)

	return &result
}

type slicePtrRange[T any] []T

func (r *slicePtrRange[T]) Empty() bool            { return r == nil || len(*r) == 0 }
func (r *slicePtrRange[T]) Front() *T              { return &(*r)[0] }
func (r *slicePtrRange[T]) PopFront()              { *r = (*r)[1:] }
func (r *slicePtrRange[T]) Save() ForwardRange[*T] { return SlicePtrRange([]T(*r)) }

// SlicePtrRange creates a range of pointers to values from a slice.
func SlicePtrRange[T any](slice []T) ForwardRange[*T] {
	result := slicePtrRange[T](slice)

	return &result
}

type sliceRetroRange[T any] []T

func (r *sliceRetroRange[T]) Empty() bool           { return r == nil || len(*r) == 0 }
func (r *sliceRetroRange[T]) Front() T              { return (*r)[len(*r)-1] }
func (r *sliceRetroRange[T]) PopFront()             { *r = (*r)[:len(*r)-1] }
func (r *sliceRetroRange[T]) Save() ForwardRange[T] { return SliceRetroRange([]T(*r)) }

// SliceRetroRange creates a range from a slice in reverse.
func SliceRetroRange[T any](slice []T) ForwardRange[T] {
	result := sliceRetroRange[T](slice)

	return &result
}

type slicePtrRetroRange[T any] []T

func (r *slicePtrRetroRange[T]) Empty() bool            { return r == nil || len(*r) == 0 }
func (r *slicePtrRetroRange[T]) Front() *T              { return &(*r)[len(*r)-1] }
func (r *slicePtrRetroRange[T]) PopFront()              { *r = (*r)[:len(*r)-1] }
func (r *slicePtrRetroRange[T]) Save() ForwardRange[*T] { return SlicePtrRetroRange([]T(*r)) }

// SlicePtrRetroRange creates a range of pointers to values from a slice.
func SlicePtrRetroRange[T any](slice []T) ForwardRange[*T] {
	result := slicePtrRetroRange[T](slice)

	return &result
}

// Slice creates a slice of memory from a range.
func Slice[T any](r InputRange[T]) []T {
	slice := make([]T, 0)

	for !r.Empty() {
		slice = append(slice, r.Front())
		r.PopFront()
	}

	return slice
}

// SliceF is `Slice` accepting a ForwardRange.
func SliceF[T any](r ForwardRange[T]) []T {
	return Slice[T](r)
}

// Bytes produces a range over the bytes of string.
func Bytes(s string) ForwardRange[byte] {
	return SliceRange([]byte(s))
}

// Runes produces a range over the runes of string.
func Runes(s string) ForwardRange[rune] {
	return SliceRange([]rune(s))
}

// String creates a new string from a range of runes.
func String(r InputRange[rune]) string {
	return string(Slice(r))
}

// StringF is `String` accepting a ForwardRange.
func StringF(r ForwardRange[rune]) string {
	return string(Slice(I(r)))
}

// StringS is `String` accepting a slice of runes.
//
// This can be used as a callback where `string` cannot be.
func StringS(r []rune) string {
	return string(r)
}
