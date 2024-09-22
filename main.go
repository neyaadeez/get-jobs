package main

import (
	"fmt"

	"github.com/neyaadeez/go-get-jobs/workday"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func main() {
	resp, err := workdaymain.GetWorkdayJobs(workday.Barclays)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp)
	fmt.Println(len(resp))
}
