package common

import (
	"fmt"
	"os"
)

// companies with own job portals
var (
	Google     = "GOGL"
	Microsoft  = "MISF"
	Oracle     = "ORCL"
	Apple      = "APLE"
	Meta       = "META"
	Tesla      = "TSLA"
	Chime      = "CHME"
	Splunk     = "SPNK"
	Visa       = "VISA"
	Uber       = "UBER"
	Databricks = "DTBS"
) //////////////////////// Edit here

var AllCompanies = make(map[string]bool)

var SitesCompanies = make(map[string]bool)

func checkDuplicatesComapnies() {
	values := []string{
		Google,
		Microsoft,
		Oracle,
		Apple,
		Meta,
		Tesla,
		Chime,
		Splunk,
		Visa,
		Uber,
		Databricks,
	} ///////////////////////// Edit here

	for _, value := range values {
		if AllCompanies[value] || SitesCompanies[value] {
			fmt.Printf("Duplicate company code found: %s\n", value)
			os.Exit(1)
		} else {
			AllCompanies[value] = true
			SitesCompanies[value] = true
		}
	}
}
