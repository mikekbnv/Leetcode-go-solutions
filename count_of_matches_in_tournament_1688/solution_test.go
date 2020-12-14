package count_of_matches_in_tournament_1688

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countMatches(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "count matches",
			args: args{
				n: 7,
			},
			want: 6,
		},
		{
			name: "count matches",
			args: args{
				n: 14,
			},
			want: 13,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, numberOfMatches(tt.args.n))
		})
	}
}
