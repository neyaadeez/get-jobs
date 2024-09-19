package sites

// Job represents the structure to hold job listing data
type JobPosting struct {
	JobTitle     string `json:"title"`
	Location     string `json:"locationsText"`
	PostedOn     string `json:"postedOn"`
	ExternalPath string `json:"externalPath"`
}

// JobsResponse represents the structure of the full response
type JobsResponse struct {
	JobPostings []JobPosting `json:"jobPostings"`
	Total       int          `json:"total"`
}
