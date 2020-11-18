package running_sum_of_1d_array_1480

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_runningSum(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "running sum of 1d array",
			args: args{
				arr: []int{
					3, 7, 10, 20, 56,
				},
			},
			want: []int{
				3, 10, 20, 40, 96,
			},
		},
		{
			name: "running sum of 1d array",
			args: args{
				arr: []int{
					3, 7, 10, 20, 56,
				},
			},
			want: []int{
				3, 10, 20, 40, 96,
			},
		},
		{
			name: "running sum of 1d array",
			args: args{
				arr: []int{
					34, 1, 15, 27, 55, 99, 100,
				},
			},
			want: []int{
				34, 35, 50, 77, 132, 231, 331,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runningSum(tt.args.arr)
			assert.Equal(t, tt.want, tt.args.arr)
		})
	}
}
