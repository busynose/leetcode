package leetcode

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		input  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test twosum",
			args: args{
				input:  []int{1, 2, 3, 4, 5, 6},
				target: 9,
			},
			want: []int{4, 3},
		},
		{
			name: "test twosum",
			args: args{
				input:  []int{1, 2, 3, 4, 5, 6},
				target: 11,
			},
			want: []int{5, 4},
		},
		{
			name: "test twosum",
			args: args{
				input:  []int{1, 1},
				target: 2,
			},
			want: []int{1, 0},
		},
		{
			name: "test twosum",
			args: args{
				input:  []int{1},
				target: 2,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.input, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
