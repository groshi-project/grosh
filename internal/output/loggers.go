package output

import (
	"log"
	"os"
)

var ErrorLogger = log.New(os.Stderr, "error: ", 0)
var StdoutLogger = log.New(os.Stdout, "", 0)

var PlusLogger = log.New(os.Stdout, "(+) ", 0)
var MinusLogger = log.New(os.Stdout, "(-) ", 0)

var TipLogger = log.New(os.Stdout, "(@) ", 0)
