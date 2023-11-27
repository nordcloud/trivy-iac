package functions

import (
	"fmt"
	"strconv"
	"strings"
)

func PadLeft(args ...interface{}) interface{} {
	if len(args) != 3 {
		return ""
	}

	input := fmt.Sprintf("%v", args[0])

	length, err := strconv.Atoi(fmt.Sprintf("%v", args[1]))
	if err != nil {
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
