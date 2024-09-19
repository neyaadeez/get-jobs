package sites

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func GetIntelJobs(days int) (JobsResponse, error) {
	client := GetClient()
	url := "https://intel.wd1.myworkdayjobs.com/External?locations=1e4a4eb3adf101b8aec18a77bf810dd0&locations=1e4a4eb3adf1018c4bf78f77bf8112d0"

	_, err := client.R().Get(url)
	if err != nil {
		log.Fatalf("Error accessing the URL: %v", err)
	}

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
		"offset": 0,
		"searchText": ""
	}`

	jobURL := "https://intel.wd1.myworkdayjobs.com/wday/cxs/intel/External/jobs"

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetHeader("X-Calypso-CSRF-Token", "YOUR_CSRF_TOKEN").
		SetBody(payload).
		Post(jobURL)

	if err != nil {
		log.Fatalf("Error fetching job listings: %v", err)
	}

	var jobsResponse JobsResponse
	err = json.Unmarshal(resp.Body(), &jobsResponse)
	if err != nil {
		log.Fatalf("Error parsing response: %v", err)
	}

	filteredJobs := filterJobsByDays(jobsResponse.JobPostings, days)

	fmt.Printf("Total Jobs: %d\n", len(filteredJobs))
	for i, job := range filteredJobs {
		filteredJobs[i].ExternalPath = "https://intel.wd1.myworkdayjobs.com/en-US/External" + job.ExternalPath
	}

	jobsResponse.JobPostings = filteredJobs
	jobsResponse.Total = len(filteredJobs)

	return jobsResponse, nil
}

func filterJobsByDays(jobPostings []JobPosting, days int) []JobPosting {
	var filteredJobs []JobPosting

	for _, job := range jobPostings {
		daysSincePosted := parsePostedOn(job.PostedOn)
		if daysSincePosted <= days {
			filteredJobs = append(filteredJobs, job)
		}
	}

	return filteredJobs
}

// Helper function to parse the "Posted On" string and return the number of days since posting
func parsePostedOn(postedOn string) int {
	lowerPostedOn := strings.ToLower(postedOn)
	if strings.Contains(lowerPostedOn, "today") {
		return 0
	} else if strings.Contains(lowerPostedOn, "yesterday") {
		return 1
	} else {
		parts := strings.Split(lowerPostedOn, " ")
		if len(parts) > 1 {
			daysAgo, err := strconv.Atoi(parts[1])
			if err == nil {
				return daysAgo
			}
		}
	}

	return 1000
}
