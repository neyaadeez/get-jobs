package sites

import (
	"encoding/json"
	"fmt"
	"strings"
)

func GetIntelJobs() ([]JobPosting, error) {

	offset := 20
	var jobPostings []JobPosting
	resp, err := intelJobs(0)
	if err != nil {
		return []JobPosting{}, err
	}

	jobPostings = append(jobPostings, resp.JobPostings...)
	for {
		if len(jobPostings) < resp.Total {
			r, err := intelJobs(offset)
			if err != nil {
				return jobPostings, err
			}

			jobPostings = append(jobPostings, r.JobPostings...)
			offset += 20
		} else {
			break
		}
	}

	return jobPostings, nil
}

func intelJobs(offset int) (JobsResponse, error) {
	client := GetClient()

	preURL := "https://intel.wd1.myworkdayjobs.com/en-US/External"
	jobURL := "https://intel.wd1.myworkdayjobs.com/wday/cxs/intel/External/jobs"
	payload := `{
		"appliedFacets": {
			"locations": [
				"1e4a4eb3adf101b8aec18a77bf810dd0",
				"1e4a4eb3adf1018c4bf78f77bf8112d0",
				"1e4a4eb3adf1013ddb7bd877bf8153d0",
				"1e4a4eb3adf10129d05fe377bf815dd0",
				"1e4a4eb3adf10118b1dfe877bf8162d0",
				"1e4a4eb3adf10155d1cc0778bf8180d0",
				"1e4a4eb3adf101d4e5a61779bf8159d1",
				"1e4a4eb3adf10146fd5c5276bf81eece",
				"1e4a4eb3adf1011246675c76bf81f8ce",
				"1e4a4eb3adf1016541777876bf8111cf",
				"1e4a4eb3adf101fa2a777d76bf8116cf",
				"1e4a4eb3adf101770f350977bf8193cf",
				"1e4a4eb3adf10174f0548376bf811bcf",
				"1e4a4eb3adf101cc4e292078bf8199d0"
			]
		},
		"limit": 20,
		"offset": %d,
		"searchText": ""
	}`

	payload = fmt.Sprintf(payload, offset)

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("X-Calypso-CSRF-Token", "YOUR_CSRF_TOKEN").
		SetBody(payload).
		Post(jobURL)

	if err != nil {
		return JobsResponse{}, fmt.Errorf("error fetching job listings: %v", err)
	}

	var jobsResponse JobsResponse
	err = json.Unmarshal(resp.Body(), &jobsResponse)
	if err != nil {
		return JobsResponse{}, fmt.Errorf("error parsing response: %v", err)
	}

	for i, job := range jobsResponse.JobPostings {
		jobsResponse.JobPostings[i].ExternalPath = preURL + job.ExternalPath
		jobid := strings.Split(jobsResponse.JobPostings[i].ExternalPath, "_")
		jobsResponse.JobPostings[i].JobId = jobid[len(jobid)-1]
	}

	return jobsResponse, nil
}
