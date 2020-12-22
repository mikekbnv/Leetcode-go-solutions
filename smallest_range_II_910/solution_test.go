package smallest_range_II_910

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SmallestRange(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "smallest range",
			args: args{
				arr: []int{
					1,
				},
				k: 0,
			},
			want: 0,
		},
		{
			name: "smallest range",
			args: args{
				arr: []int{
					0, 10,
				},
				k: 2,
			},
			want: 6,
		},
		{
			name: "smallest range",
			args: args{
				arr: []int{
					1, 3, 6,
				},
				k: 3,
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, smallestRangeII(tt.args.arr, tt.args.k))
		})
	}
}
