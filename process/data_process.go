package process

import (
	"fmt"
	"sync"

	"github.com/neyaadeez/go-get-jobs/common"
	"github.com/neyaadeez/go-get-jobs/database"
	"github.com/neyaadeez/go-get-jobs/sites"
	"github.com/neyaadeez/go-get-jobs/workday"
)

var (
	cachedJobs  []common.JobPosting
	cachedError error
	onceGetJobs sync.Once
)

func ProcessJobsWithDB() error {
	jobs, err := GetProcessedNewJobs()
	if err != nil {
		fmt.Println("error while processing new jobs: ", err.Error())
		return err
	}
	fmt.Println("Processed Jobs (New Jobs): ", len(jobs))

	err = database.InsertIntoDB(jobs)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func GetProcessedNewJobs() ([]common.JobPosting, error) {
	onceGetJobs.Do(func() {
		var allJobs []common.JobPosting

		jobs, err := GetAllWorkdayJobs()
		if err != nil {
			fmt.Println(err.Error())
			cachedError = err
			return
		}
		fmt.Println("All Workday Jobs: ", len(jobs))
		allJobs = append(allJobs, jobs...)

		jobs, err = sites.GetGoogleJobs()
		if err != nil {
			fmt.Println(err.Error())
			cachedError = err
			return
		}
		fmt.Println("All Google Jobs: ", len(jobs))
		allJobs = append(allJobs, jobs...)

		jobs, err = sites.GetMicrosoftJobs()
		if err != nil {
			fmt.Println(err.Error())
			cachedError = err
			return
		}
		fmt.Println("All Microsoft Jobs: ", len(jobs))
		allJobs = append(allJobs, jobs...)

		jobs, err = sites.GetOracleJobs()
		if err != nil {
			fmt.Println(err.Error())
			cachedError = err
			return
		}
		fmt.Println("All Oracle Jobs: ", len(jobs))
		allJobs = append(allJobs, jobs...)

		jobs, err = sites.GetAppleJobs()
		if err != nil {
			fmt.Println(err.Error())
			cachedError = err
			return
		}
		fmt.Println("All Apple Jobs: ", len(jobs))
		allJobs = append(allJobs, jobs...)

		cachedJobs, cachedError = processDublicateJobs(allJobs)
		if cachedError != nil {
			fmt.Println(cachedError.Error())
		}
	})

	return cachedJobs, cachedError
}

func ProcessJobsWithDBForNewlyAddedJobPortal() error {
	workday.Init()
	jobs, err := getProcessedNewJobsNewlyAddedJobPortal()
	if err != nil {
		fmt.Println("error while processing new jobs: ", err.Error())
		return err
	}
	fmt.Println("Processed Jobs (New Jobs): ", len(jobs))

	err = database.InsertIntoDB(jobs)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func getProcessedNewJobsNewlyAddedJobPortal() ([]common.JobPosting, error) {
	onceGetJobs.Do(func() {
		var allJobs []common.JobPosting

		jobs, err := sites.GetMetaJobs()
		if err != nil {
			fmt.Println(err.Error())
			cachedError = err
			return
		}
		fmt.Println("All Apple Jobs: ", len(jobs))
		allJobs = append(allJobs, jobs...)

		// jobs, err := workdaymain.GetWorkdayJobs(workdaymain.WorkdayPayloads[common.Tancent])
		// if err != nil {
		// 	fmt.Println(err.Error())
		// 	cachedError = err
		// 	return
		// }
		// fmt.Println(jobs[0])
		// fmt.Println("All Tancent Jobs: ", len(jobs))
		// allJobs = append(allJobs, jobs...)

		cachedJobs, cachedError = processDublicateJobs(allJobs)
		if cachedError != nil {
			fmt.Println(cachedError.Error())
		}
	})

	return cachedJobs, cachedError
}
