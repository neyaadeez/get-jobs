package sites

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/neyaadeez/go-get-jobs/common"
)

func GetGoogleJobs() (common.JobsResponse, error) {
	client := common.GetClient()

	url := "https://careers.google.com/jobs/results/?location=United%20States&target_level=EARLY&target_level=INTERN_AND_APPRENTICE&sort_by=date"

	resp, err := client.R().Get(url)
	if err != nil {
		return common.JobsResponse{}, fmt.Errorf("error accessing the URL: %v", err)
	}

	return parseJobPage(resp.Body())
}

// Function to parse the HTML content using goquery
func parseJobPage(body []byte) (common.JobsResponse, error) {
	doc, err := goquery.NewDocumentFromReader(io.NopCloser(bytes.NewReader(body)))
	if err != nil {
		return common.JobsResponse{}, fmt.Errorf("error parsing the HTML: %v", err)
	}

	var jobs []common.JobPosting
	count := 0

	doc.Find(".WpHeLc").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")

		ariaLabel, _ := s.Attr("aria-label")

		jobs = append(jobs, common.JobPosting{
			JobTitle:     strings.ReplaceAll(ariaLabel, "Learn more about", ""),
			ExternalPath: "https://www.google.com/about/careers/applications/" + href,
		})
		count++
	})

	return common.JobsResponse{JobPostings: jobs, Total: count}, nil
}
