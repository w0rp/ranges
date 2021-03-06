package ranges

import "testing"

func TestZipN(t *testing.T) {
	t.Parallel()

	output := Slice(
		Map(
			ZipN(
				Only[any](1, 2, 3),
				Only[any](4, 5, 6),
				Only[any](7, 8, 9),
			),
			func(e []any) []int {
				return SliceF(
					MapF(
						SliceRange(e),
						func(x any) int { return x.(int) },
					),
				)
			},
		),
	)

	assertEqual(t, output, [][]int{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	})
}
