package workday

import (
	"github.com/neyaadeez/go-get-jobs/common"
	workdaymain "github.com/neyaadeez/go-get-jobs/workday_main"
)

func init() {
	workdaymain.RegisterPayload(common.Barclays, common.WorkdayPayload{
		CmpCode: common.Barclays,
		PreURL:  "https://barclays.wd3.myworkdayjobs.com/en-US/External_Career_Site_Barclays",
		JobsURL: "https://barclays.wd3.myworkdayjobs.com/wday/cxs/barclays/External_Career_Site_Barclays/jobs",
		PayLoad: `{
  "appliedFacets": {
    "workerSubType": [
      "6139d325cdcc1001a72ce8fbe2290000",
      "6139d325cdcc1001a72ceb63d5d60000",
      "6139d325cdcc1001a72ceb63d5d60001"
    ],
    "jobFamilyGroup": [
      "1ab48a98eb7c1001e8e0bdc7d4a10000",
      "1ab48a98eb7c1001e8e0c2947eeb0000",
      "112c054282011001e9161cb8b7960000",
      "112c054282011001e9162cfccdc10000"
    ],
    "timeType": [
      "259ef6e735f8101411549dcf4d1e0003"
    ]
  },
  "limit": 20,
  "offset": %d,
  "searchText": ""
}`,
	})
}
