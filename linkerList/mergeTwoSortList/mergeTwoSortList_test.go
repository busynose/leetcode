package leetcode

import (
	"testing"
)

func Test_mergeTwoSortList(t *testing.T) {

	l1 := buildNodeList([]int{1, 2, 3, 4, 5})
	l2 := buildNodeList([]int{1, 1, 3, 4, 5})
	l3 := buildNodeList([]int{1, 1, 1, 2, 3, 3, 4, 4, 5, 5})
	type args struct {
		l1 *Node
		l2 *Node
	}

	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "mergetwoList",
			args: args{
				l1: l1,
				l2: l2,
			},
			want: l3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoSortList(tt.args.l1, tt.args.l2); !isnodeListEqual(got, tt.want) {
				t.Errorf("mergeTwoSortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
