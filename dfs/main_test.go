package dfs

import "testing"

func Test_minimumTimeRequired(t *testing.T) {
	type args struct {
		jobs []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "jobs = [3,2,3], k = 3",
			args: args{
				jobs: []int{3,2,3},
				k:    3,
			},
			want: 3,
		},
		{
			name: "jobs = [1,2,4,7,8], k = 2",
			args: args{
				jobs: []int{1,2,4,7,8},
				k:    2,
			},
			want: 11,
		},
		{
			name: "jobs = [8,6,3,3,2,2], k=2",
			args: args{
				jobs: []int{8,6,3,3,2,2},
				k:    2,
			},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTimeRequired(tt.args.jobs, tt.args.k); got != tt.want {
				t.Errorf("minimumTimeRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}
