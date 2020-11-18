package merge_intervals_56

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func Test_merge(t *testing.T) {
	type args struct {
		arr [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "merge intervals",
			args: args{
				arr: [][]int{
					{1, 3},
					{2, 6},
					{8, 10},
					{15, 18},
				},
			},
			want: [][]int{
					{1, 6},
					{8, 10},
					{15, 18},
			},
		},
		{
			name: "merge intervals",
			args: args{
				arr: [][]int{
					{1, 4},
					{4, 5},
				},
			},
			want: [][]int{
					{1, 5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := merge(tt.args.arr)
			assert.Equal(t, tt.want, res)
		})
	}
}