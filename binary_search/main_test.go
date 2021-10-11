package binary_search

import (
	"testing"
)

func Test_findInMountainArray(t *testing.T) {
	type args struct {
		target      int
		mountainArr *MountainArray
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "array = [1,2,3,4,5,3,1], target = 3",
			args: args{
				target:      3,
				mountainArr: NewMountainArray([]int{1,2,3,4,5,3,1}),
			},
			want: 2,
		},
		{
			name: "array = [0,1,2,4,2,1], target = 3",
			args: args{
				target:      3,
				mountainArr: NewMountainArray([]int{0,1,2,4,2,1}),
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findInMountainArray(tt.args.target, tt.args.mountainArr); got != tt.want {
				t.Errorf("findInMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
