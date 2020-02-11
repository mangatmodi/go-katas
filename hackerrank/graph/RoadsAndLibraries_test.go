package graph

import "testing"

func Test_roadsAndLibraries(t *testing.T) {
	type args struct {
		n      int32
		c_lib  int32
		c_road int32
		cities [][]int32
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := roadsAndLibraries(tt.args.n, tt.args.c_lib, tt.args.c_road, tt.args.cities); got != tt.want {
				t.Errorf("roadsAndLibraries() = %v, want %v", got, tt.want)
			}
		})
	}
}