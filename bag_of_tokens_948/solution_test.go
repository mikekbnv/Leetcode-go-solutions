package bag_of_tokens_948

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BagOfTokens(t *testing.T) {
	type args struct {
		arr []int
		P int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "bag of tokens",
			args: args{
				arr: []int{
					100, 
				},
				P: 50,
			},
			want: 0,
		},
		{
			name: "bag of tokens",
			args: args{
				arr: []int{
					100, 200,
				},
				P: 150,
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, bagOfTokensScore(tt.args.arr, tt.args.P))
		})
	}
}
