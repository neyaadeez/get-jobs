package common

// Job represents the structure to hold job listing data
type JobPosting struct {
	JobId        string
	JobTitle     string `json:"title"`
	Location     string `json:"locationsText,omitempty"`
	PostedOn     string `json:"postedOn,omitempty"`
	ExternalPath string `json:"externalPath"`
}

// JobsResponse represents the structure of the full response
type JobsResponse struct {
	JobPostings []JobPosting `json:"jobPostings"`
	Total       int          `json:"total"`
}

// WorkdayPayload
type WorkdayPayload struct {
	CmpCode string
	PreURL  string
	JobsURL string
	PayLoad string
}
