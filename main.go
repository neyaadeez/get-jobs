package main

import (
	"fmt"

	"github.com/neyaadeez/go-get-jobs/process"
)

func main() {
	resp, err := process.GetAllWorkdayJobs()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp)
	fmt.Println(len(resp))
}
