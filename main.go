package main

import (
	"fmt"

	"github.com/neyaadeez/go-get-jobs/database"
	"github.com/neyaadeez/go-get-jobs/process"
)

func main() {
	resp, err := process.GetAllWorkdayJobs()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("total New Jobs Posted Today: ", len(resp))

	err = database.InsertIntoDB(resp)
	if err != nil {
		fmt.Println(err.Error())
	}
}
