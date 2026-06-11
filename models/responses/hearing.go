package responses

import (
	"litigation_backend/utils"
	"time"
)

type Hearing struct {
	Id        *string    `json:"id" firestore:"id"`
	CaseId    string     `json:"case_id" firestore:"case_id"`
	Date      time.Time  `json:"date" firestore:"date"`
	Submitted string     `json:"submitted" firestore:"submitted"`
	Happened  string     `json:"happened" firestore:"happened"`
	Order     string     `json:"order" firestore:"order"`
	NextDate  *time.Time `json:"next_date" firestore:"next_date"`
}

func HearingFromMap(data map[string]any) Hearing {
	return Hearing{
		Id:        utils.HandleNil[string](data["id"]),
		CaseId:    data["case_id"].(string),
		Date:      data["date"].(time.Time),
		Submitted: data["submitted"].(string),
		Happened:  data["happened"].(string),
		Order:     data["order"].(string),
		NextDate:  utils.HandleNil[time.Time](data["next_date"]),
	}
}
