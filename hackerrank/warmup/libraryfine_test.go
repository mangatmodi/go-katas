package warmup

import "testing"

func Test_libraryFine(t *testing.T) {
	type args struct {
		d1 int32
		m1 int32
		y1 int32
		d2 int32
		m2 int32
		y2 int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{"days late", args{9, 6, 2015, 6, 6, 2015}, 45},
		{"years late", args{9, 6, 2016, 6, 6, 2015}, 10000},
		{"failing case", args{2, 7, 1014, 1, 1, 1015}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := libraryFine(tt.args.d1, tt.args.m1, tt.args.y1, tt.args.d2, tt.args.m2, tt.args.y2); got != tt.want {
				t.Errorf("libraryFine() = %v, want %v", got, tt.want)
			}
		})
	}
}
