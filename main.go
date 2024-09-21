package main

import (
	"fmt"

	"github.com/neyaadeez/go-get-jobs/workday"
)

func main() {
	resp, err := workday.GetWorkdayJobs(workday.SalesForce)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp)
	fmt.Println(len(resp))
}
