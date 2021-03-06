package ranges

import "testing"

func TestNull(t *testing.T) {
	t.Parallel()

	null := Null[int]()

	assertEmptyF(t, null)
}

func TestOnly(t *testing.T) {
	t.Parallel()

	assertEqual(t, SliceF(Only[int]()), []int{})
	assertEqual(t, SliceF(Only(1)), []int{1})
	assertEqual(t, SliceF(Only(1, 2)), []int{1, 2})
	assertEqual(t, SliceF(Only(1, 2, 3)), []int{1, 2, 3})
}
