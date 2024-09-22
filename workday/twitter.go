package workday

import "github.com/neyaadeez/go-get-jobs/common"

var Twitter = common.WorkdayPayload{
	CmpCode: common.Twitter,
	PreURL:  "https://twitter.wd5.myworkdayjobs.com/en-US/X",
	JobsURL: "https://twitter.wd5.myworkdayjobs.com/wday/cxs/twitter/X/jobs",
	PayLoad: `{
  "appliedFacets": {},
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
}
