package bulls_and_cows_299

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BullsandCows(t *testing.T) {
	type args struct {
		secret string
		guess  string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "bulls and cows",
			args: args{
				secret: "1807",
				guess:  "7810",
			},
			want: "1A3B",
		},
		{
			name: "bulls and cows",
			args: args{
				secret: "1123",
				guess:  "0111",
			},
			want: "1A1B",
		},
		{
			name: "bulls and cows",
			args: args{
				secret: "1",
				guess:  "0",
			},
			want: "0A0B",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, getHint(tt.args.secret, tt.args.guess))
		})
	}
}
