package wzap

import (
	"fmt"
	"testing"
)

func Test_ParseLevel(t *testing.T) {
	lvstr := "INFO|Debu"
	lv := parseLevel(lvstr, "|")
	if lv == InfoLevel|DebugLevel {
		fmt.Println("true")
	}
}
