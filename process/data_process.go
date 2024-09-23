package process

import (
	"fmt"

	"github.com/neyaadeez/go-get-jobs/common"
	"github.com/neyaadeez/go-get-jobs/database"
	"github.com/neyaadeez/go-get-jobs/sites"
)

func ProcessJobsWithDB() error {
	var allJobs []common.JobPosting
	jobs, err := GetAllWorkdayJobs()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("All Workday Jobs: ", len(jobs))

	allJobs = append(allJobs, jobs...)

	jobs, err = sites.GetGoogleJobs()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("All Google Jobs: ", len(jobs))

	allJobs = append(allJobs, jobs...)

	jobs, err = sites.GetMicrosoftJobs()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("All Microsoft Jobs: ", len(jobs))

	allJobs = append(allJobs, jobs...)

	// process Jobs
	jobs, err = processDublicateJobs(allJobs)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Processed Jobs(New Jobs): ", len(jobs))

	err = database.InsertIntoDB(jobs)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
