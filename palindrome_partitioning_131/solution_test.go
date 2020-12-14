package palindrome_partitioning_131

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_palindromePartitioning(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "palindrome partitioning",
			args: args{
				s: "aab",
			},
			want: [][]string{
				{"a", "a", "b"},
				{"aa", "b"},
			},
		},
		{
			name: "palindrome partitioning",
			args: args{
				s: "a",
			},
			want: [][]string{
				{"a"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, partition(tt.args.s))
		})
	}
}
