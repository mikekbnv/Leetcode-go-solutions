package goal_parser_interpretation_1678

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_goalParser(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "goal parser",
			args: args{
				str: "G()(al)",
			},
			want: "Goal",
		},
		{
			name: "goal parser",
			args: args{
				str: "G()()()()(al)",
			},
			want: "Gooooal",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, interpret(tt.args.str))
		})
	}
}
