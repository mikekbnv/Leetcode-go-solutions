package increasing_triplet_subsequence_334

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_tripletSubsequence(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "triplet subsequence",
			args: args{
				nums: []int{
					1, 2, 3, 4, 5,
				},
			},
			want: true,
		},
		{
			name: "triplet subsequence",
			args: args{
				nums: []int{
					5, 4, 3, 2, 1,
				},
			},
			want: false,
		},
		{
			name: "triplet subsequence",
			args: args{
				nums: []int{
					2, 1, 5, 0, 4, 6,
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, increasingTriplet(tt.args.nums))
		})
	}
}
