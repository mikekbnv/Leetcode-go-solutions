package top_k_frequent_words_692

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_topKfrequentWords(t *testing.T) {
	type args struct {
		arr []string
		k int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "top k frequent words",
			args: args{
				arr: []string{
					"i", "love", "leetcode", "i", "love", "coding",
				},
				k: 2,
			},
			want: []string{
				"i", "love",
			},
		},
		{
			name: "top k frequent words",
			args: args{
				arr: []string{
					"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is",
				},
				k: 4,
			},
			want: []string{
				"the", "is", "sunny", "day",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, topKFrequent(tt.args.arr, tt.args.k))
		})
	}
}
