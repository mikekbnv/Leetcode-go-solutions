package unique_morse_code_words_804

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_runningSum(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "unique morse code words",
			args: args{
				arr: []string{
					"gin", "zen", "gig", "msg",
				},
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, uniqueMorseRepresentations(tt.args.arr))
		})
	}
}
