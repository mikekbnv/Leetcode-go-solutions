package valid_anagrama_242

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validAnagrama(t *testing.T) {
	type args struct {
		str1 string
		str2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid anagrama",
			args: args{
				str1: "anagram",
				str2: "nagaram",
			},
			want: true,
		},
		{
			name: "valid anagrama",
			args: args{
				str1: "rat",
				str2: "car",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isAnagram(tt.args.str1, tt.args.str2))
		})
	}
}
