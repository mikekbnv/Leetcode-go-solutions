package find_first_and_last_position_of_element_in_sorted_array_34

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findfirstandlast(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "find first and last position of element in sorted array",
			args: args{
				arr: []int{
					5, 7, 7, 8, 8, 10,
				},
				target: 8,
			},
			want: []int{
				3, 4,
			},
		},
		{
			name: "find first and last position of element in sorted array",
			args: args{
				arr: []int{
					5, 7, 7, 8, 8, 10,
				},
				target: 6,
			},
			want: []int{
				-1, -1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, searchRange(tt.args.arr, tt.args.target))
		})
	}
}
