package edit_distance_72

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_editDistance(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "edit distance",
			args: args{
				word1: "horse",
				word2: "ros",
			},
			want: 3,
		},
		{
			name: "edit distance",
			args: args{
				word1: "intention",
				word2: "execution",
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, minDistance(tt.args.word1, tt.args.word2))
		})
	}
}
