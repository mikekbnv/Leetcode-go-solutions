package squares_of_a_sorted_array_977

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_squaresArray(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "squares array",
			args: args{
				arr: []int{
					-4, -1, 0, 3, 10,
				},
			},
			want: []int{
				0, 1, 9, 16, 100,
			},
		},
		{
			name: "squares array",
			args: args{
				arr: []int{
					-7, -3, 2, 3, 11,
				},
			},
			want: []int{
				4, 9, 9, 49, 121,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, sortedSquares(tt.args.arr))
		})
	}
}
