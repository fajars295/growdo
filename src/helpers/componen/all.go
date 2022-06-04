package componen

import (
	"fmt"
	"strconv"
)

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	fmt.Println(err)
	return i
}
func StringToBool(s string) bool {
	i, err := strconv.ParseBool(s)
	fmt.Println(err)
	return i
}

func SwitchCase(s string) string {
	switch s {
	case "blog":
		return "type"
	default:
		return ""
	}
}
