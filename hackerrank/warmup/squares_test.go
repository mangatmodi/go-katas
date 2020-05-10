package warmup

import "testing"

func Test_squares(t *testing.T) {
	type args struct {
		a int32
		b int32
	}
	tests := []struct {
		name string
		a    int32
		b    int32
		want int32
	}{
		{"base case", 24, 49, 3},
		{"base case", 17, 24, 0},
		{"base case", 3, 9, 2},
		{"base case", 3, 9, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := squares(tt.a, tt.b); got != tt.want {
				t.Errorf("squares() = %v, want %v", got, tt.want)
			}
		})
	}
}
