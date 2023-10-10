package output

import (
	"log"
	"os"
)

var Error = log.New(os.Stderr, "error: ", 0)
var Stdout = log.New(os.Stdout, "", 0)

var Plus = log.New(os.Stdout, "(+) ", 0)
var Minus = log.New(os.Stdout, "(-) ", 0)

var Tip = log.New(os.Stdout, "(@) ", 0)
