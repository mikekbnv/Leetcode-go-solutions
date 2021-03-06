package search_in_rotated_sorted_array_33

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchInArray(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "search in rotated sorted array",
			args: args{
				arr: []int{
					4, 5, 6, 7, 0, 1, 2,
				},
				target: 0,
			},
			want: 4,
		},
		{
			name: "search in rotated sorted array",
			args: args{
				arr: []int{
					4, 5, 6, 7, 0, 1, 2,
				},
				target: 3,
			},
			want: -1,
		},
		{
			name: "search in rotated sorted array",
			args: args{
				arr: []int{
					1,
				},
				target: 0,
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, search(tt.args.arr, tt.args.target))
		})
	}
}
