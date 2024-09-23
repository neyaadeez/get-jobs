package main

import (
	"fmt"

	"github.com/neyaadeez/go-get-jobs/sites"
)

func main() {

	// err := process.ProcessJobsWithDB()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// workday.Init()
	resp, err := sites.GetOracleJobs()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp)
	fmt.Println(len(resp))
}
