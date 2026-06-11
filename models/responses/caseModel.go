package responses

import (
	"litigation_backend/utils"
	"time"
)

func CaseModelFromJson(data map[string]any) CaseModel {
	plaintiffs := make([]string, 0)
	if data["plaintiffs"] != nil {
		for _, plaintiff := range data["plaintiffs"].([]any) {
			plaintiffs = append(plaintiffs, plaintiff.(string))
		}
	}
	respondents := make([]string, 0)
	if data["respondents"] != nil {
		for _, respondent := range data["respondents"].([]any) {
			respondents = append(respondents, respondent.(string))
		}
	}
	return CaseModel{
		Id:           utils.HandleNil[string](data["id"]),
		UserId:       utils.HandleNil[string](data["user_id"]),
		CaseNo:       data["case_no"].(string),
		Year:         int(data["year"].(int64)),
		Court:        Court(data["court"].(string)),
		Bench:        Bench(data["bench"].(string)),
		Title:        data["title"].(string),
		Plaintiffs:   plaintiffs,
		Respondents:  respondents,
		FirstHearing: utils.ToTimePtr(data["first_hearing"]),
		LastHearing:  utils.ToTimePtr(data["last_hearing"]),
		NextHearing:  utils.ToTimePtr(data["next_hearing"]),
		Status:       CaseStatus(data["status"].(string)),
		Notes:        data["notes"].(string),
		CaseNature:   utils.ToConst[CaseNature](data["case_nature"]),
		Department:   utils.ToConst[Department](data["department"]),
		Taluka:       utils.ToConst[Taluka](data["taluka"]),
		Documents:    toAppDocumentSlice(data["documents"].([]any)),
		Hearings:     make([]Hearing, 0),
	}
}

type CaseModel struct {
	Id           *string       `json:"id" firestore:"id"`
	UserId       *string       `json:"user_id" firestore:"user_id"`
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
	CaseNature   *CaseNature   `json:"case_nature" firestore:"case_nature"`
	Department   *Department   `json:"department" firestore:"department"`
	Taluka       *Taluka       `json:"taluka" firestore:"taluka"`
	Documents    []AppDocument `json:"documents" firestore:"documents"`
	Hearings     []Hearing     `json:"hearings"`
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

func toAppDocumentSlice(data []any) []AppDocument {
	documents := make([]AppDocument, 0)
	for _, doc := range data {
		docMap := doc.(map[string]any)
		document := AppDocument{
			Id:         docMap["id"].(string),
			Type:       utils.ToConst[DocumentType](docMap["type"]),
			Name:       docMap["name"].(string),
			Url:        docMap["file_name"].(string),
			UploadedAt: docMap["uploaded_at"].(time.Time),
		}
		documents = append(documents, document)
	}
	return documents
}
