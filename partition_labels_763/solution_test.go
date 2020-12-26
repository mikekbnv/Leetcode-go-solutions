package partition_labels_763

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PartitionLables(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "partition lables",
			args: args{
				s: "ababcbacadefegdehijhklij",
			},
			want: []int{
				9, 7, 8,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, partitionLabels(tt.args.s))
		})
	}
}
