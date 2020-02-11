package warmup

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestDiagonalDifference_Load(t *testing.T) {
	type args struct {
		stdin io.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantStdout string
	}{
		{
			"Simple case",
			args{strings.NewReader(`3
11 2 4
4 5 6
10 8 -12
			`)},
			"15\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			di := DiagonalDifference{}
			stdout := &bytes.Buffer{}
			di.Load(tt.args.stdin, stdout)
			if gotStdout := stdout.String(); gotStdout != tt.wantStdout {
				t.Errorf("Load() = %v, want %v", gotStdout, tt.wantStdout)
			}
		})
	}
}
