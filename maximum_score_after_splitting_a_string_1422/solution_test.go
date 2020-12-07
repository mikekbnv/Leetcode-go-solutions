package maximum_score_after_splitting_a_string_1422

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximumScore(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "maximum score after splitting a string",
			args: args{
				s: "011101",
			},
			want: 5,
		},
		{
			name: "maximum score after splitting a string",
			args: args{
				s: "00111",
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxScore(tt.args.s))
		})
	}
}
