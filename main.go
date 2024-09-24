package main

import (
	"fmt"

	"github.com/neyaadeez/go-get-jobs/process"
	"github.com/neyaadeez/go-get-jobs/readme"
)

func main() {

	err := process.ProcessJobsWithDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = readme.ReadMeProcessNewJobs()
	if err != nil {
		fmt.Println("error while processing readme file with new jobs: ")
	}

	// // workday.Init()
	// resp, err := sites.GetOracleJobs()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println(resp)
	// fmt.Println(len(resp))
}
