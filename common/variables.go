package common

import (
	"fmt"
	"os"
)

func init() {
	checkDuplicates()
}

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

func checkDuplicates() {
	valueMap := make(map[string]bool)

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
	}

	for _, value := range values {
		if valueMap[value] {
			fmt.Printf("Duplicate company code found: %s\n", value)
			os.Exit(1)
		} else {
			valueMap[value] = true
		}
	}
}
