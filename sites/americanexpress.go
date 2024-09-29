package sites

type AmexMain struct {
	Positions []AmexPosition `json:"positions"`
	Count     int64          `json:"count"`
}

type AmexPosition struct {
	ID                   int64       `json:"id"`
	Name                 string      `json:"name"`
	Location             string      `json:"location"`
	Locations            []string    `json:"locations"`
	Hot                  int64       `json:"hot"`
	Department           string      `json:"department"`
	TUpdate              int64       `json:"t_update"`
	TCreate              int64       `json:"t_create"`
	AtsJobID             string      `json:"ats_job_id"`
	DisplayJobID         string      `json:"display_job_id"`
	IDLocale             string      `json:"id_locale"`
	JobDescription       string      `json:"job_description"`
	Stars                int64       `json:"stars"`
	MedallionProgram     interface{} `json:"medallionProgram"`
	LocationFlexibility  interface{} `json:"location_flexibility"`
	CanonicalPositionURL string      `json:"canonicalPositionUrl"`
	IsPrivate            bool        `json:"isPrivate"`
}
