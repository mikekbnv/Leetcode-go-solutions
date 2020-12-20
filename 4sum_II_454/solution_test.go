package four_sum_II_454

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_4Sum(t *testing.T) {
	type args struct {
		A []int
		B []int
		C []int
		D []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "4 sum",
			args: args{
				A: []int{
					1, 2,
				},
				B: []int{
					-2, -1,
				},
				C: []int{
					-1, 2,
				},
				D: []int{
					0, 2,
				},
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, fourSumCount(tt.args.A, tt.args.B, tt.args.C, tt.args.D))
		})
	}
}
