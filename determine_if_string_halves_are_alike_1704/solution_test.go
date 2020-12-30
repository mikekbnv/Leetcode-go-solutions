package determine_if_string_halves_are_alike_1704

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DetermineString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "determine string",
			args: args{
				s: "book",
			},
			want: true,
		},
		{
			name: "determine string",
			args: args{
				s: "textbook",
			},
			want: false,
		},
		{
			name: "determine string",
			args: args{
				s: "MerryChristmas",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, halvesAreAlike(tt.args.s))
		})
	}
}
