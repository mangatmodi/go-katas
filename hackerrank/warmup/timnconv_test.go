package warmup

import "testing"

func Test_timeConversion(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"11 am", "11:59:59AM", "11:59:59"},
		{"12 am", "12:59:59AM", "00:59:59"},
		{"12 pm", "12:59:59PM", "12:59:59"},
		{"3 pm", "3:59:59PM", "15:59:59"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeConversion(tt.s); got != tt.want {
				t.Errorf("timeConversion() = %v, want %v", got, tt.want)
			}
		})
	}
}
