package gas_station_134

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_gasStation(t *testing.T) {
	type args struct {
		arr1 []int
		arr2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "gas station",
			args: args{
				arr1: []int{
					1, 2, 3, 4, 5,
				},
				arr2: []int{
					3, 4, 5, 1, 2,
				},
			},
			want: 3,
		},
		{
			name: "gas station",
			args: args{
				arr1: []int{
					2, 3, 4,
				},
				arr2: []int{
					3, 4, 3,
				},
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, canCompleteCircuit(tt.args.arr1, tt.args.arr2))
		})
	}
}
