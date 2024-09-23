package main

import (
	"fmt"

	"github.com/neyaadeez/go-get-jobs/process"
)

func main() {

	err := process.ProcessJobsWithDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// workday.Init()
	// resp, err := workdaymain.GetWorkdayJobs(workdaymain.WorkdayPayloads["BARC"])
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println(resp)
	// fmt.Println(len(resp))
}
