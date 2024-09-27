package common

import (
	"fmt"
	"os"
)

// companies with own job portals
var (
	Google    = "GOGL"
	Microsoft = "MISF"
	Oracle    = "ORCL"
	Apple     = "APLE"
	Meta      = "META"
	Tesla     = "TSLA"
	Chime     = "CHME"
)

var Companies = make(map[string]bool)

func checkDuplicatesComapnies() {
	values := []string{
		Google,
		Microsoft,
		Oracle,
		Apple,
		Meta,
		Tesla,
	}

	for _, value := range values {
		if Companies[value] {
			fmt.Printf("Duplicate company code found: %s\n", value)
			os.Exit(1)
		} else {
			Companies[value] = true
		}
	}
}
