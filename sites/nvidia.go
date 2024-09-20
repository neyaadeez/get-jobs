package sites

import (
	"encoding/json"
	"fmt"
	"strings"
)

func GetNvidiaJobs() ([]JobPosting, error) {

	offset := 20
	var jobPostings []JobPosting
	resp, err := nvidiaJobs(0)
	if err != nil {
		return []JobPosting{}, err
	}

	jobPostings = append(jobPostings, resp.JobPostings...)
	for {
		if len(jobPostings) < resp.Total {
			r, err := nvidiaJobs(offset)
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

func nvidiaJobs(offset int) (JobsResponse, error) {
	client := GetClient()

	preURL := "https://nvidia.wd5.myworkdayjobs.com/en-US/NVIDIAExternalCareerSite"
	jobURL := "https://nvidia.wd5.myworkdayjobs.com/wday/cxs/nvidia/NVIDIAExternalCareerSite/jobs"
	payload := `{
  "appliedFacets": {
    "locationHierarchy1": [
      "2fcb99c455831013ea52fb338f2932d8"
    ],
    "jobFamilyGroup": [
      "0c40f6bd1d8f10ae43ffaefd46dc7e78",
      "0c40f6bd1d8f10ae43ffc3fc7d8c7e8a",
      "0c40f6bd1d8f10ae43ffc668c6847e8c",
      "0c40f6bd1d8f10ae43ffbd1459047e84"
    ],
    "workerSubType": [
      "ab40a98049581037a3ada55b087049b7"
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
