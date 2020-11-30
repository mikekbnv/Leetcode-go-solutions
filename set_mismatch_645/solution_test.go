package set_mismatch_645

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_setMismatch(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "set mismatch",
			args: args{
				arr: []int{
					1, 2, 2, 4,
				},
			},
			want: []int{
				2, 3,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findErrorNums(tt.args.arr))
		})
	}
}
