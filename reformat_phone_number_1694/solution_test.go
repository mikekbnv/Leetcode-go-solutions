package reformat_phone_number_1694

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReformatNumber(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "reformat number",
			args: args{
				num: "1-23-45 6",
			},
			want: "123-456",
		},
		{
			name: "reformat number",
			args: args{
				num: "123 4-567",
			},
			want: "123-45-67",
		},
		{
			name: "reformat number",
			args: args{
				num: "123 4-5678",
			},
			want: "123-456-78",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, reformatNumber(tt.args.num))
		})
	}
}
