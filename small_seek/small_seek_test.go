package small_seek

import (
	"reflect"
	"testing"
)

func TestSmallerNumnersThanCurrent(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[8, 1, 2, 2, 3]",
			args: args{nums: []int{8, 1, 2, 2, 3}},
			want: []int{4, 0, 1, 1, 3},
		},
		{
			name: "[6, 5, 4, 8]",
			args: args{nums: []int{6, 5, 4, 8}},
			want: []int{2, 1, 0, 3},
		},
		{
			name: "[7, 7, 7, 7]",
			args: args{nums: []int{7, 7, 7, 7}},
			want: []int{0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SmallerNumnersThanCurrent(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SmallerNumnersThanCurrent() = %v, want %v", got, tt.want)
			}
		})
	}
}
