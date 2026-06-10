package responses

import "time"

func CaseModelFromJson(data map[string]any) CaseModel {
	return CaseModel{
		Id:           data["id"].(string),
		UserId:       data["user_id"].(string),
		CaseNo:       data["case_no"].(string),
		Year:         int(data["year"].(int64)),
		Court:        Court(data["court"].(string)),
		Bench:        Bench(data["bench"].(string)),
		Title:        data["title"].(string),
		Plaintiffs:   data["plaintiffs"].([]string),
		Respondents:  data["respondents"].([]string),
		FirstHearing: toTimePtr(data["first_hearing"]),
		LastHearing:  toTimePtr(data["last_hearing"]),
		NextHearing:  toTimePtr(data["next_hearing"]),
		Status:       CaseStatus(data["status"].(string)),
		Notes:        data["notes"].(string),
		CaseNature:   CaseNature(data["case_nature"].(string)),
		Department:   Department(data["department"].(string)),
		Taluka:       Taluka(data["taluka"].(string)),
		Documents:    toAppDocumentSlice(data["documents"].([]any)),
	}
}

type CaseModel struct {
	Id           string        `json:"id" firestore:"id"`
	UserId       string        `json:"user_id" firestore:"user_id"`
	CaseNo       string        `json:"case_no" firestore:"case_no"`
	Year         int           `json:"year" firestore:"year"`
	Court        Court         `json:"court" firestore:"court"`
	Bench        Bench         `json:"bench" firestore:"bench"`
	Title        string        `json:"title" firestore:"title"`
	Plaintiffs   []string      `json:"plaintiffs" firestore:"plaintiffs"`
	Respondents  []string      `json:"respondents" firestore:"respondents"`
	FirstHearing *time.Time    `json:"first_hearing" firestore:"first_hearing"`
	LastHearing  *time.Time    `json:"last_hearing" firestore:"last_hearing"`
	NextHearing  *time.Time    `json:"next_hearing" firestore:"next_hearing"`
	Status       CaseStatus    `json:"status" firestore:"status"`
	Notes        string        `json:"notes" firestore:"notes"`
	CaseNature   CaseNature    `json:"case_nature" firestore:"case_nature"`
	Department   Department    `json:"department" firestore:"department"`
	Taluka       Taluka        `json:"taluka" firestore:"taluka"`
	Documents    []AppDocument `json:"documents" firestore:"documents"`
}

type Court string

const (
	highCourt           Court = "highCourt"
	civilCourt          Court = "civilCourt"
	sessionsCourt       Court = "sessionsCourt"
	antiCorruptionCourt Court = "antiCorruptionCourt"
	servicesTribunal    Court = "servicesTribunal"
	federalShariatCourt Court = "federalShariatCourt"
	supremeCourt        Court = "supremeCourt"
	otherCourt          Court = "other"
)

type Bench string

const (
	singleBench   Bench = "singleBench"
	divisionBench Bench = "divisionBench"
	fullBench     Bench = "fullBench"
)

type CaseNature string

const (
	revision              CaseNature = "revision"
	generalRecruitment    CaseNature = "generalRecruitment"
	disabledQuota         CaseNature = "disabledQuota"
	generalAdministration CaseNature = "generalAdministration"
	serviceMatter         CaseNature = "serviceMatter"
	contemptCase          CaseNature = "contemptCase"
)

type Taluka string

const (
	sukkurCity Taluka = "sukkurCity"
	newSukkur  Taluka = "newSukkur"
	rohri      Taluka = "rohri"
	panoAqil   Taluka = "panoAqil"
	salehPath  Taluka = "salehPath"
)

type Department string

const (
	educationLiteracy Department = "educationLiteracy"
	health            Department = "health"
	irrigation        Department = "irrigation"
	agriculture       Department = "agriculture"
	revenue           Department = "revenue"
	policeHome        Department = "policeHome"
	localGovernment   Department = "localGovernment"
	worksServices     Department = "worksServices"
	socialWelfare     Department = "socialWelfare"
	otherDepartment   Department = "other"
)

type CaseStatus string

const (
	active      CaseStatus = "active"
	stayGranted CaseStatus = "stayGranted"
	decided     CaseStatus = "decided"
	dismissed   CaseStatus = "dismissed"
)

func toTimePtr(t any) *time.Time {
	if t == nil {
		return nil
	}
	timeValue := t.(time.Time)
	return &timeValue
}

func toAppDocumentSlice(data []any) []AppDocument {
	documents := make([]AppDocument, 0)
	for _, doc := range data {
		docMap := doc.(map[string]any)
		document := AppDocument{
			Id:         docMap["id"].(string),
			Type:       DocumentType(docMap["type"].(string)),
			Name:       docMap["name"].(string),
			Url:        docMap["url"].(string),
			UploadedAt: docMap["uploaded_at"].(time.Time),
		}
		documents = append(documents, document)
	}
	return documents
}
