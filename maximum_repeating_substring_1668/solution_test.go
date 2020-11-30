package maximum_repeating_substring_1668

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MaximumRepeatingSub(t *testing.T) {
	type args struct {
		seq  string
		word string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "maximum repeating substring",
			args: args{
				seq:  "ababc",
				word: "ab",
			},
			want: 2,
		},
		{
			name: "maximum repeating substring",
			args: args{
				seq:  "ababc",
				word: "ba",
			},
			want: 1,
		},
		{
			name: "maximum repeating substring",
			args: args{
				seq:  "ababc",
				word: "ac",
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxRepeating(tt.args.seq, tt.args.word))
		})
	}
}
