// package process

// import (
// 	"fmt"

// 	"github.com/neyaadeez/go-get-jobs/common"
// 	"github.com/neyaadeez/go-get-jobs/workday"
// 	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
// )

// func GetAllWorkdayJobs() ([]common.JobPosting, error) {
// 	workday.Init()
// 	var allJobs []common.JobPosting

// 	for company := range common.WorkdayCompanies {
// 		payload, ok := workdaymain.WorkdayPayloads[company]
// 		if !ok {
// 			fmt.Printf("No Workday payload found for company: %s\n", company)
// 			continue
// 		}

// 		jobs, err := workdaymain.GetWorkdayJobs(payload)
// 		if err != nil {
// 			fmt.Printf("Error fetching jobs for %s: %v\n", company, err)
// 			continue
// 		}

// 		allJobs = append(allJobs, jobs...)
// 	}

// 	return allJobs, nil
// }

package process

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/neyaadeez/go-get-jobs/common"
	"github.com/neyaadeez/go-get-jobs/workday"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

const jobIDFile = "job_ids.json"

func loadJobIDs() (map[string]struct{}, error) {
	file, err := os.Open(jobIDFile)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]struct{}), nil
		}
		return nil, err
	}
	defer file.Close()

	var jobIDs []string
	if err := json.NewDecoder(file).Decode(&jobIDs); err != nil {
		return nil, err
	}

	jobIDSet := make(map[string]struct{})
	for _, id := range jobIDs {
		jobIDSet[id] = struct{}{}
	}
	return jobIDSet, nil
}

func saveJobIDs(jobIDSet map[string]struct{}) error {
	jobIDs := make([]string, 0, len(jobIDSet))
	for id := range jobIDSet {
		jobIDs = append(jobIDs, id)
	}

	data, err := json.Marshal(jobIDs)
	if err != nil {
		return err
	}
	return os.WriteFile(jobIDFile, data, 0644)
}

func GetAllWorkdayJobs() ([]common.JobPosting, error) {
	workday.Init()
	var allJobs []common.JobPosting
	jobIDSet, err := loadJobIDs()
	if err != nil {
		return nil, err
	}

	for company := range common.WorkdayCompanies {
		payload, ok := workdaymain.WorkdayPayloads[company]
		if !ok {
			fmt.Printf("No Workday payload found for company: %s\n", company)
			continue
		}

		jobs, err := workdaymain.GetWorkdayJobs(payload)
		if err != nil {
			fmt.Printf("Error fetching jobs for %s: %v\n", company, err)
			continue
		}

		for _, job := range jobs {
			if _, exists := jobIDSet[job.JobId]; !exists {
				allJobs = append(allJobs, job)
				jobIDSet[job.JobId] = struct{}{}
			}
		}
	}

	if err := saveJobIDs(jobIDSet); err != nil {
		return nil, err
	}

	return allJobs, nil
}
