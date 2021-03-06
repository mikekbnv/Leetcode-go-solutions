package kth_missing_positive_number_1539

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_kth_missing(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "kth missing positive number",
			args: args{
				arr: []int{
					2, 3, 4, 7, 11,
				},
				k: 5,
			},
			want: 9,
		},
		{
			name: "kth missing positive number",
			args: args{
				arr: []int{
					1, 2, 3, 4,
				},
				k: 2,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findKthPositive(tt.args.arr, tt.args.k))
		})
	}
}
