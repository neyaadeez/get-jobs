package sites

import (
	"encoding/json"
	"fmt"
	"strings"
)

func GetCrowdStrikeJobs() ([]JobPosting, error) {

	offset := 20
	var jobPostings []JobPosting
	resp, err := crowdStrikeJobs(0)
	if err != nil {
		return []JobPosting{}, err
	}

	jobPostings = append(jobPostings, resp.JobPostings...)
	for {
		if len(jobPostings) < resp.Total {
			r, err := crowdStrikeJobs(offset)
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

func crowdStrikeJobs(offset int) (JobsResponse, error) {
	client := GetClient()

	preURL := "https://crowdstrike.wd5.myworkdayjobs.com/en-US/crowdstrikecareers"
	jobURL := "https://crowdstrike.wd5.myworkdayjobs.com/wday/cxs/crowdstrike/crowdstrikecareers/jobs"
	payload := `{
  "appliedFacets": {
    "locationCountry": [
      "bc33aa3152ec42d4995f4791a106ed09"
    ],
    "Job_Family": [
      "cb19f044639b1001f6a02595bc920000",
      "1408861ee6e201641be2c2f6b000c00b",
      "1408861ee6e20197f95adbf6b000d20b",
      "1408861ee6e201d67af3e0f6b000d60b",
      "1408861ee6e2015adbe3e7f6b000de0b",
      "1408861ee6e201df327f0ff7b000fa0b"
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
