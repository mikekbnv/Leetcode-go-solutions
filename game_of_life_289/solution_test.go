package game_of_life_289

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_gameofLife(t *testing.T) {
	type args struct {
		board [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "game of life",
			args: args{
				board: [][]int{
					{0, 1, 0},
					{0, 0, 1},
					{1, 1, 1},
					{0, 0, 0},
				},
			},
			want: [][]int{
				{0, 0, 0},
				{1, 0, 1},
				{0, 1, 1},
				{0, 1, 0},
			},
		},
		{
			name: "game of life",
			args: args{
				board: [][]int{
					{1, 1},
					{1, 0},
				},
			},
			want: [][]int{
				{1, 1},
				{1, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, gameOfLife(tt.args.board))
		})
	}
}
