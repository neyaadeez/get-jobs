package main

import (
	"fmt"

	"github.com/neyaadeez/go-get-jobs/sites"
)

func main() {
	// resp, err := sites.GetGoogleJobs()
	// resp, err := sites.GetMicrosoftJobs(1, "")
	// resp, err := sites.GetOracleJobs(1, "")
	// resp, err := sites.GetIntelJobs()
	// resp, err := sites.GetCrowdStrikeJobs()
	// resp, err := sites.GetQualysJobs()
	resp, err := sites.GetNvidiaJobs()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(len(resp))
	fmt.Println(resp)
}
