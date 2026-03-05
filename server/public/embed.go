package public

import (
	_ "embed"
)

//go:embed index.html
var IndexHTML string

//go:embed scalar.html
var ScalarHTML []byte
