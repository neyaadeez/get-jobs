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

type OracleJob struct {
	Id                 string `json:"Id"`
	Title              string `json:"Title"`
	PostedDate         string `json:"PostedDate"`
	PrimaryLocation    string `json:"PrimaryLocation"`
	ShortDescription   string `json:"ShortDescriptionStr"`
	SecondaryLocations []struct {
		Name string `json:"Name"`
	} `json:"secondaryLocations"`
}

type OracleJobResponse struct {
	Items []struct {
		RequisitionList []OracleJob `json:"requisitionList"`
	} `json:"items"`
}

func GetOracleJobs(days int, keyword string) (common.JobsResponse, error) {
	client := common.GetClient()

	url := formatOracleURL("https://eeho.fa.us2.oraclecloud.com/hcmRestApi/resources/latest/recruitingCEJobRequisitions", keyword, days)

	// url := "https://eeho.fa.us2.oraclecloud.com/hcmRestApi/resources/latest/recruitingCEJobRequisitions?onlyData=true&expand=requisitionList.secondaryLocations,flexFieldsFacet.values,requisitionList.requisitionFlexFields&finder=findReqs;siteNumber=CX_45001,facetsList=LOCATIONS%3BWORK_LOCATIONS%3BWORKPLACE_TYPES%3BTITLES%3BCATEGORIES%3BORGANIZATIONS%3BPOSTING_DATES%3BFLEX_FIELDS,limit=50,lastSelectedFacet=POSTING_DATES,locationId=300000000149325,selectedCategoriesFacet=300000001917356%3B300000001917346,selectedLocationsFacet=300000000149325,selectedPostingDatesFacet=7,sortBy=POSTING_DATES_DESC"
	resp, err := client.R().Get(url)
	if err != nil {
		return common.JobsResponse{}, fmt.Errorf("error accessing the URL: %v", err)
	}

	var jobsResponseOracle OracleJobResponse
	err = json.Unmarshal(resp.Body(), &jobsResponseOracle)
	if err != nil {
		return common.JobsResponse{}, fmt.Errorf("error parsing response: %v", err)
	}

	filteredJobs := filterOracleJobsByDays(jobsResponseOracle.Items[0].RequisitionList, days)

	var jobPostings []common.JobPosting
	for _, job := range filteredJobs {
		jobPosting := common.JobPosting{
			JobTitle:     job.Title,
			Location:     formatOracleLocations(job.PrimaryLocation, job.SecondaryLocations),
			PostedOn:     job.PostedDate,
			ExternalPath: generateOracleJobLink(job.Id, job.Title),
		}
		jobPostings = append(jobPostings, jobPosting)
	}

	return common.JobsResponse{
		JobPostings: jobPostings,
		Total:       len(jobPostings),
	}, nil
}

func filterOracleJobsByDays(jobs []OracleJob, days int) []OracleJob {
	var filteredJobs []OracleJob

	for _, job := range jobs {
		daysSincePosted := parseOraclePostedOn(job.PostedDate)
		if daysSincePosted <= days {
			filteredJobs = append(filteredJobs, job)
		}
	}

	return filteredJobs
}

func parseOraclePostedOn(postingDate string) int {
	parsedTime, err := time.Parse("2006-01-02", postingDate) // Adjust the date format if necessary
	if err != nil {
		log.Printf("Error parsing posting date: %v", err)
		return 1000
	}

	duration := time.Since(parsedTime)
	return int(duration.Hours() / 24)
}

func formatOracleLocations(primary string, secondary []struct {
	Name string `json:"Name"`
}) string {
	var locations []string
	locations = append(locations, primary)

	for _, loc := range secondary {
		locations = append(locations, loc.Name)
	}

	return strings.Join(locations, "; ")
}

func generateOracleJobLink(jobID, jobTitle string) string {
	baseURL := "https://eeho.fa.us2.oraclecloud.com/hcmRestApi/resources/latest/recruitingCEJobRequisitions"
	encodedTitle := url.PathEscape(strings.ReplaceAll(jobTitle, " ", "-"))
	return fmt.Sprintf("%s/%s/%s", baseURL, jobID, encodedTitle)
}

func formatOracleURL(baseURL, keyword string, days int) string {
	queryParams := url.Values{}
	queryParams.Set("onlyData", "true")
	queryParams.Set("expand", "requisitionList.secondaryLocations,flexFieldsFacet.values,requisitionList.requisitionFlexFields")

	if keyword != "" {
		//keyword = url.QueryEscape(fmt.Sprintf(`"%s"`, keyword)) // This escapes the keyword with double quotes
		queryParams.Set("finder", fmt.Sprintf("findReqs;siteNumber=CX_45001,facetsList=LOCATIONS,3BWORK_LOCATIONS,3BWORKPLACE_TYPES,3BTITLES,3BCATEGORIES,3BORGANIZATIONS,3BPOSTING_DATES,3BFLEX_FIELDS,limit=50,keyword=%s,lastSelectedFacet=POSTING_DATES,locationId=300000000149325,selectedCategoriesFacet=300000001917356,3B300000001917346,selectedLocationsFacet=300000000149325,selectedPostingDatesFacet=%d,sortBy=POSTING_DATES_DESC", keyword, days))
	} else {
		queryParams.Set("finder", fmt.Sprintf("findReqs;siteNumber=CX_45001,facetsList=LOCATIONS,WORK_LOCATIONS,WORKPLACE_TYPES,TITLES,CATEGORIES,ORGANIZATIONS,POSTING_DATES,FLEX_FIELDS,limit=50,lastSelectedFacet=POSTING_DATES,locationId=300000000149325,selectedCategoriesFacet=300000001917356,300000001917346,selectedLocationsFacet=300000000149325,selectedPostingDatesFacet=%d,sortBy=POSTING_DATES_DESC", days))
	}

	return baseURL + "?" + queryParams.Encode()
}
