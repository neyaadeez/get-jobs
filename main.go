package main

import (
	"fmt"

	"github.com/neyaadeez/go-get-jobs/process"
	"github.com/neyaadeez/go-get-jobs/readme"
	"github.com/neyaadeez/go-get-jobs/sites"
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
	//processTodaysJobsDBAndReadme()
	//process.ProcessJobsWithDBForNewlyAddedJobPortal(common.Snowflake, false)

	// workday.Init()
	resp, err := sites.GetSnowflakeJobs()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(resp)
	fmt.Println(len(resp))

	// fmt.Println(resp[0])
	// fmt.Println(len(resp))

	// jobs, err := workdaymain.GetWorkdayJobs(workdaymain.WorkdayPayloads[common.Walmart])
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println(jobs[0])
	// fmt.Println("All Jobs: ", len(jobs))

}
