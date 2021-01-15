package word_search_79

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_wordsearch(t *testing.T) {
	type args struct {
		board [][]string
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "word search",
			args: args{
				board: [][]string{
					{"A", "B", "C", "E"},
					{"S", "F", "C", "S"},
					{"A", "D", "E", "E"},
				},
				word: "ABCCED",
			},
			want: true,
		},
		{
			name: "word search",
			args: args{
				board: [][]string{
					{"A", "B", "C", "E"},
					{"S", "F", "C", "S"},
					{"A", "D", "E", "E"},
				},
				word: "SEE",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, wordsearch(tt.args.board, tt.args.word))
		})
	}
}
