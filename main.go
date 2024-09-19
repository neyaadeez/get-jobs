package main

import (
	"fmt"

	"github.com/neyaadeez/go-get-jobs/sites"
)

func main() {
	resp, err := sites.GetIntelJobs(4)
	if err != nil {
		return
	}

	fmt.Println(resp)
}
