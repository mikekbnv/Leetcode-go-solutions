package container_with_most_water

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_4Sum(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Container With Most Water",
			args: args{
				height: []int{
					1, 8, 6, 2, 5, 4, 8, 3, 7,
				},
			},
			want: 49,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxArea(tt.args.height))
		})
	}
}
