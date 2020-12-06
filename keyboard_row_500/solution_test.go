package keyboard_row_500

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_keybordRow(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "keybord row",
			args: args{
				arr: []string{
					"Hello", "Alaska", "Dad", "Peace",
				},
			},
			want: []string{
				"Alaska", "Dad",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findWords(tt.args.arr))
		})
	}
}
