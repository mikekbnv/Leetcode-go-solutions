package sliding_window_maximum_239

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SlidingWindow(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "sliding window maximum",
			args: args{
				arr: []int{
					1, 3, -1, -3, 5, 3, 6, 7,
				},
				k: 3,
			},
			want: []int{
				3, 3, 5, 5, 6, 7,
			},
		},
		{
			name: "sliding window maximum",
			args: args{
				arr: []int{
					1,
				},
				k: 1,
			},
			want: []int{
				1,
			},
		},
		{
			name: "sliding window maximum",
			args: args{
				arr: []int{
					4, -2,
				},
				k: 2,
			},
			want: []int{
				4,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxSlidingWindow(tt.args.arr, tt.args.k))
		})
	}
}
