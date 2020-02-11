package hackerrank

import "io"

type Solver interface {
	Load(stdin io.Reader, stdout io.Writer)
}
