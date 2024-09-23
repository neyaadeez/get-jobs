package common

import (
	"fmt"
	"os"
)

func init() {
	checkDuplicatesWorkday()
}

// WorkDay companies list
var (
	ASML                    = "ASML"
	CrowdStrike             = "CRWD"
	Intel                   = "INTL"
	Nvidia                  = "NVDA"
	Qualys                  = "QLYS"
	SalesForce              = "SLSF"
	Walmart                 = "WLMT"
	Target                  = "TRGT"
	Samsung                 = "SMSN"
	Disney                  = "DSNY"
	Sony                    = "SONY"
	Twitter                 = "TWTR"
	CapitalOne              = "CONE"
	Boeing                  = "BOEG"
	Bose                    = "BOSE"
	Snapchat                = "SNPT"
	CVS                     = "CVSS"
	CCCIntelligentSolutions = "CCCI"
	NorthropGrumman         = "NTGN"
	Phinia                  = "PHNA"
	Nissan                  = "NISN"
	HP                      = "HPHP"
	Barclays                = "BARC"
)

var WorkdayCompanies = make(map[string]bool)

func checkDuplicatesWorkday() {
	values := []string{
		ASML,
		CrowdStrike,
		Intel,
		Nvidia,
		Qualys,
		SalesForce,
		Walmart,
		Target,
		Samsung,
		Disney,
		Sony,
		Twitter,
		CapitalOne,
		Boeing,
		Bose,
		Snapchat,
		CVS,
		CCCIntelligentSolutions,
		NorthropGrumman,
		Phinia,
		Nissan,
		HP,
		Barclays,
	}

	for _, value := range values {
		if WorkdayCompanies[value] {
			fmt.Printf("Duplicate company code found: %s\n", value)
			os.Exit(1)
		} else {
			WorkdayCompanies[value] = true
		}
	}
}
