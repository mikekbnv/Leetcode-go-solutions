package count_servers_that_communicate_1267

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countServersThatCommunicate(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "count servers",
			args: args{
				grid: [][]int{
					{1, 0},
					{0, 1},
				},
			},
			want: 0,
		},
		{
			name: "count servers",
			args: args{
				grid: [][]int{
					{1, 0},
					{1, 1},
				},
			},
			want: 3,
		},
		{
			name: "count servers",
			args: args{
				grid: [][]int{
					{1, 1, 0, 0},
					{0, 0, 1, 0},
					{0, 0, 1, 0},
					{0, 0, 0, 1},
				},
			},
			want: 4,
		},
		
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, countServers(tt.args.grid))
		})
	}
}
