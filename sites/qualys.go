package sites

import (
	"encoding/json"
	"fmt"
	"strings"
)

func GetQualysJobs() ([]JobPosting, error) {

	offset := 20
	var jobPostings []JobPosting
	resp, err := qualysJobs(0)
	if err != nil {
		return []JobPosting{}, err
	}

	jobPostings = append(jobPostings, resp.JobPostings...)
	for {
		if len(jobPostings) < resp.Total {
			r, err := qualysJobs(offset)
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

func qualysJobs(offset int) (JobsResponse, error) {
	client := GetClient()

	preURL := "https://qualys.wd5.myworkdayjobs.com/en-US/Careers"
	jobURL := "https://qualys.wd5.myworkdayjobs.com/wday/cxs/qualys/Careers/jobs"
	payload := `{
  "appliedFacets": {
    "locations": [
      "f92a9a956d98018ceda9b0f0f5cd6ecb",
      "f92a9a956d9801d8f97abaf0f5cd78cb",
      "f92a9a956d9801c51769b5f0f5cd73cb",
      "f92a9a956d98011784169af0f5cd55cb",
      "f92a9a956d98014963aa9ef0f5cd5acb",
      "f92a9a956d98018f451da3f0f5cd5fcb",
      "f92a9a956d98010bee83a7f0f5cd64cb",
      "f92a9a956d9801d060a8cbf0f5cd8ccb",
      "f92a9a956d98012ee343d0f0f5cd91cb",
      "f92a9a956d9801f0c11bd5f0f5cd96cb",
      "f92a9a956d98016820ede3f0f5cda5cb",
      "f92a9a956d9801e96024e8f0f5cdaacb",
      "f92a9a956d98010c7972f1f0f5cdb4cb",
      "f92a9a956d980180fa31faf0f5cdbecb",
      "f92a9a956d9801f7d4b607f1f5cdcdcb",
      "f92a9a956d9801a8536210f1f5cdd7cb",
      "f92a9a956d980175e5c740f1f5cd0ecc",
      "f92a9a956d980193ba7c3cf1f5cd09cc",
      "f92a9a956d9801fd49bc49f1f5cd18cc",
      "f92a9a956d9801f7968b83f1f5cd57cc",
      "f92a9a956d98017f53d54df1f5cd1dcc",
      "134325fb3bfa0191e066155fc3ca5050",
      "f92a9a956d9801f6c21852f1f5cd22cc",
      "f92a9a956d9801756bbc5af1f5cd2ccc",
      "f92a9a956d9801e48d19dcf1f5cdc2cc",
      "f92a9a956d980154135f61f1f5cd31cc",
      "f92a9a956d9801ebe77dedf1f5cdd6cc",
      "f92a9a956d9801b9054e6bf1f5cd3bcc"
    ],
    "jobFamilyGroup": [
      "e78d04040ce81001e4412df7fa1e0000",
      "efd4289a13971001e4c227e96f660000",
      "e78d04040ce81001e478bb2ee4320000",
      "e78d04040ce81001e4726b049df30000"
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
