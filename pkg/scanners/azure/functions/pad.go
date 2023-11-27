package functions

import (
	"fmt"
	"strings"
)

func PadLeft(args ...interface{}) interface{} {
	if len(args) != 3 {
		return ""
	}

	input, ok := args[0].(string)
	if !ok {
		input = fmt.Sprintf("%v", args[0])
	}

	length, ok := args[1].(int)
	if !ok {
		return ""
	}

	pad, ok := args[2].(string)
	if !ok {
		return ""
	}

	if len(input) >= length {
		return input
	}

	repeat := (length - len(input)) / len(pad)

	return strings.Repeat(pad, repeat) + input
}
