package plus_one_66

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "plus one",
			args: args{
				arr: []int{
					1, 2, 3,
				},
			},
			want: []int{
				1, 2, 4,
			},
		},
		{
			name: "plus one",
			args: args{
				arr: []int{
					4, 3, 2, 1,
				},
			},
			want: []int{
				4, 3, 2, 2,
			},
		},
		{
			name: "plus one",
			args: args{
				arr: []int{
					9, 9, 9, 9,
				},
			},
			want: []int{
				1, 0, 0, 0, 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, plusOne(tt.args.arr))
		})
	}
}
