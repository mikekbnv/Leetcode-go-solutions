package smallest_string_with_given_numeric_value_1663

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_smallestString(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "smallest string with given numeric value",
			args: args{
				n: 3,
				k: 27,
			},
			want: "aay",
		},
		{
			name: "smallest string with given numeric value",
			args: args{
				n: 5,
				k: 73,
			},
			want: "aaszz",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, getSmallestString(tt.args.n, tt.args.k))
		})
	}
}
