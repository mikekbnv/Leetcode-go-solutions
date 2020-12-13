package count_the_number_of_consistent_strings_1684

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CountConsistentStrings(t *testing.T) {
	type args struct {
		allowed string
		words   []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "count consistent strings",
			args: args{
				words: []string{
					"ad", "bd", "aaab", "baa", "badab",
				},
				allowed: "ab",
			},
			want: 2,
		},
		{
			name: "count consistent strings",
			args: args{
				words: []string{
					"a", "b", "c", "ab", "ac", "bc", "abc",
				},
				allowed: "abc",
			},
			want: 7,
		},
		{
			name: "count consistent strings",
			args: args{
				words: []string{
					"cc", "acd", "b", "ba", "bac", "bad", "ac", "d",
				},
				allowed: "cad",
			},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, countConsistentStrings(tt.args.allowed, tt.args.words))
		})
	}
}
