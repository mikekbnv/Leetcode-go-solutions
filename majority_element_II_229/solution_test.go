package majotity_elements_II_229

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Majority_element(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "majority element",
			args: args{
				arr: []int{
					3, 2, 3,
				},
			},
			want: []int{
				3,
			},
		},
		{
			name: "majority element",
			args: args{
				arr: []int{
					1, 2,
				},
			},
			want: []int{
				1, 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, majorityElement(tt.args.arr))
		})
	}
}
