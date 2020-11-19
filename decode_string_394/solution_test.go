package decode_string_394

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_runningSum(t *testing.T) {
	type args struct {
		encodedstr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "decode string",
			args: args{
				encodedstr: "3[a]2[bc]",
			},
			want: "aaabcbc",
		},
		{
			name: "decode string",
			args: args{
				encodedstr: "3[a2[c]]",
			},
			want: "accaccacc",
		},
		{
			name: "decode string",
			args: args{
				encodedstr: "abc3[cd]xyz",
			},
			want: "abccdcdcdxyz",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, decodeString(tt.args.encodedstr))
		})
	}
}
