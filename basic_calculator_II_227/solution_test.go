package basic_calculator_II_227

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calculate(t *testing.T) {
	type args struct {
		str string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "basic calculator",
			args: args{
				str: "3+2*2",
			},
			want: 7,
		},
		{
			name: "basic calculator",
			args: args{
				str: "3/2",
			},
			want: 1,
		},
		{
			name: "basic calculator",
			args: args{
				str: "3+5 / 2",
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, calculate(tt.args.str))
		})
	}
}
