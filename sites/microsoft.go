package sites

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/neyaadeez/go-get-jobs/common"
)

type MicrosoftJob struct {
	JobID       string `json:"jobId"`
	Title       string `json:"title"`
	PostingDate string `json:"postingDate"`
	Properties  struct {
		Locations       []string `json:"locations"`
		PrimaryLocation string   `json:"primaryLocation"`
		EmploymentType  string   `json:"employmentType"`
	} `json:"properties"`
}

type MicrosoftJobResponse struct {
	OperationResult struct {
		Result struct {
			Jobs []MicrosoftJob `json:"jobs"`
		} `json:"result"`
	} `json:"operationResult"`
}

func GetMicrosoftJobs(days int, keyword string) (common.JobsResponse, error) {
	client := common.GetClient()

	url := formatURL("https://gcsservices.careers.microsoft.com/search/api/v1/search", keyword)

	// := "https://gcsservices.careers.microsoft.com/search/api/v1/search?q=software&lc=United%20States&exp=Students%20and%20graduates&l=en_us&pg=1&pgSz=20&o=Recent&flt=true"

	resp, err := client.R().Get(url)
	if err != nil {
		return common.JobsResponse{}, fmt.Errorf("error accessing the URL: %v", err)
	}

	var jobsResponseMicrosoft MicrosoftJobResponse
	err = json.Unmarshal(resp.Body(), &jobsResponseMicrosoft)
	if err != nil {
		return common.JobsResponse{}, fmt.Errorf("error parsing response: %v", err)
	}

	filteredJobs := filterMicrosoftJobsByDays(jobsResponseMicrosoft.OperationResult.Result.Jobs, days)

	var jobPostings []common.JobPosting
	for _, job := range filteredJobs {
		jobPosting := common.JobPosting{
			JobTitle:     job.Title,
			Location:     formatLocations(job.Properties.Locations),
			PostedOn:     job.PostingDate,
			ExternalPath: generateMicrosoftJobLink(job.JobID, job.Title),
		}
		jobPostings = append(jobPostings, jobPosting)
	}

	return common.JobsResponse{
		JobPostings: jobPostings,
		Total:       len(jobPostings),
	}, nil
}

func filterMicrosoftJobsByDays(jobs []MicrosoftJob, days int) []MicrosoftJob {
	var filteredJobs []MicrosoftJob

	for _, job := range jobs {
		daysSincePosted := parseMicrosoftPostedOn(job.PostingDate)
		if daysSincePosted <= days {
			filteredJobs = append(filteredJobs, job)
		}
	}

	return filteredJobs
}

func parseMicrosoftPostedOn(postingDate string) int {
	parsedTime, err := time.Parse(time.RFC3339, postingDate)
	if err != nil {
		log.Printf("Error parsing posting date: %v", err)
		return 1000
	}

	duration := time.Since(parsedTime)
	return int(duration.Hours() / 24)
}

func formatLocations(locations []string) string {
	if len(locations) == 0 {
		return "Unknown"
	}

	location := strings.Join(locations, "; ")
	return location
}

// generateMicrosoftJobLink dynamically creates the job link using job ID and title
func generateMicrosoftJobLink(jobID, jobTitle string) string {
	baseURL := "https://jobs.careers.microsoft.com/global/en/job"
	encodedTitle := url.PathEscape(strings.ReplaceAll(jobTitle, " ", "-"))
	return fmt.Sprintf("%s/%s/%s", baseURL, jobID, encodedTitle)
}

func formatURL(baseURL, keyword string) string {
	queryParams := url.Values{}
	queryParams.Set("q", keyword)
	queryParams.Set("lc", "United States")
	queryParams.Set("exp", "Students and graduates")
	queryParams.Set("l", "en_us")
	queryParams.Set("pg", "1")
	queryParams.Set("pgSz", "20")
	queryParams.Set("o", "Recent")
	queryParams.Set("flt", "true")

	return baseURL + "?" + queryParams.Encode()
}
