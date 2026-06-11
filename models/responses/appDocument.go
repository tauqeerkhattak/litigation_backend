package responses

import "time"

type AppDocument struct {
	Id         string        `json:"id" firestore:"id"`
	Type       *DocumentType `json:"type" firestore:"type"`
	Name       string        `json:"name" firestore:"name"`
	Url        string        `json:"url" firestore:"url"`
	UploadedAt time.Time     `json:"uploaded_at" firestore:"uploaded_at"`
}

type DocumentType string

const (
	plaint            DocumentType = "plaint"
	parawiseComments  DocumentType = "parawiseComments"
	writtenStatement  DocumentType = "writtenStatement"
	complianceReport  DocumentType = "complianceReport"
	orderSheet        DocumentType = "orderSheet"
	judgement         DocumentType = "judgment"
	otherDocumentType DocumentType = "other"
)
