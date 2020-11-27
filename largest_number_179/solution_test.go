package largest_number_179

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_latgestnumber(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "largest number",
			args: args{
				arr: []int{
					10, 2,
				},
			},
			want: "210",
		},
		{
			name: "largest number",
			args: args{
				arr: []int{
					3,30,34,5,9,
				},
			},
			want: "9534330",
		},
		{
			name: "largest number",
			args: args{
				arr: []int{
					1,
				},
			},
			want: "1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, largestNumber(tt.args.arr))
		})
	}
}
