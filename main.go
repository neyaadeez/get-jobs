package main

import (
	"fmt"

	"github.com/neyaadeez/go-get-jobs/process"
	"github.com/neyaadeez/go-get-jobs/readme"
)

func processTodaysJobsDBAndReadme() {
	err := process.ProcessJobsWithDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = readme.ReadMeProcessNewJobs()
	if err != nil {
		fmt.Println("error while processing readme file with new jobs: ")
	}
}

func main() {
	processTodaysJobsDBAndReadme()
	//process.ProcessJobsWithDBForNewlyAddedJobPortal(common.Databricks, false)

	// workday.Init()
	// resp, err := sites.GetDatabricksJobs()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println(resp[0])
	// fmt.Println(len(resp))

}
