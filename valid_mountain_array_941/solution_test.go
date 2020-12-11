package valid_mountain_array_941

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validMountain(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid mauntain",
			args: args{
				arr: []int{
					2, 1,
				},
			},
			want: false,
		},
		{
			name: "valid mauntain",
			args: args{
				arr: []int{
					3, 5, 5,
				},
			},
			want: false,
		},
		{
			name: "valid mauntain",
			args: args{
				arr: []int{
					0, 3, 2, 1,
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, validMountainArray(tt.args.arr))
		})
	}
}
