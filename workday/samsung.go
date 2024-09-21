package workday

import "github.com/neyaadeez/go-get-jobs/common"

var Samsung = common.WorkdayPayload{
	CmpCode: common.Samsung,
	PreURL:  "https://sec.wd3.myworkdayjobs.com/en-US/Samsung_Careers",
	JobsURL: "https://sec.wd3.myworkdayjobs.com/wday/cxs/sec/Samsung_Careers/jobs",
	PayLoad: `{
  "appliedFacets": {
    "Location_Country": [
      "bc33aa3152ec42d4995f4791a106ed09"
    ]
  },
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
}
