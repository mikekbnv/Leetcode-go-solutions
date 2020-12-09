package pairs_of_songs_with_total_durations_divisible_by_60_1010

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pairOfsongs(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "pair of songs",
			args: args{
				arr: []int{
					30, 20, 150, 100, 40,
				},
			},
			want: 3,
		},
		{
			name: "pair of songs",
			args: args{
				arr: []int{
					60, 60, 60,
				},
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, numPairsDivisibleBy60(tt.args.arr))
		})
	}
}
