package sort_transformed_array

import (
	"reflect"
	"testing"
)

func Test_sortTransformedArray(t *testing.T) {
	type args struct {
		nums []int
		a    int
		b    int
		c    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "nums = [-4,-2,2,4], a = 1, b = 3, c = 5",
			args: args{
				nums: []int{-4,-2,2,4},
				a:    1,
				b:    3,
				c:    5,
			},
			want: []int{3,9,15,33},
		},
		{
			name: "nums = [-4,-2,2,4], a = -1, b = 3, c = 5",
			args: args{
				nums: []int{-4,-2,2,4},
				a:    -1,
				b:    3,
				c:    5,
			},
			want: []int{-23,-5,1,7},
		},
		{
			name: "nums=[-4,-2,2,4], a=0, b=3, c=5",
			args: args{
				nums: []int{-4,-2,2,4},
				a:    0,
				b:    3,
				c:    5,
			},
			want: []int{-7, -1, 11, 17},
		},
		{
			name: "nums = [-4,-2,2,4], a=0, b=-3, c=5",
			args: args{
				nums: []int{-4,-2,2,4},
				a:    0,
				b:    -3,
				c:    5,
			},
			want: []int{-7, -1, 11, 17},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortTransformedArray(tt.args.nums, tt.args.a, tt.args.b, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortTransformedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
