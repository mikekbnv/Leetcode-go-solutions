package count_of_smaller_numbers_after_self_315

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Count_of_Smaller_Numbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "count of smaller numbers",
			args: args{
				nums: []int {
					5, 2, 6, 1,
				},
			},
			want: []int{
				2, 1, 1, 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, countSmaller(tt.args.nums))
		})
	}
}
