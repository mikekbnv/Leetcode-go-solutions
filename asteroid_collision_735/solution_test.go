package asteroid_collision_735

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_asteroidColl(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "asteroid collision",
			args: args{
				arr: []int{
					5, 10, -5,
				},
			},
			want: []int{
				5, 10,
			},
		},
		{
			name: "asteroid collision",
			args: args{
				arr: []int{
					8, -8,
				},
			},
			want: []int{},
		},
		{
			name: "asteroid collision",
			args: args{
				arr: []int{
					10, 2, -5,
				},
			},
			want: []int{
				10,
			},
		},
		{
			name: "asteroid collision",
			args: args{
				arr: []int{
					-2, -1, 1, 2,
				},
			},
			want: []int{
				-2, -1, 1, 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := asteroidCollision(tt.args.arr)
			assert.Equal(t, tt.want, res)
		})
	}
}
