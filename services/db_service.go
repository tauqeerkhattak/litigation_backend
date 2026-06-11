package services

import (
	"context"
	"litigation_backend/config"
	"litigation_backend/models/requests"
	"litigation_backend/models/responses"
	"time"

	"cloud.google.com/go/firestore/apiv1/firestorepb"
)

func CreateUserInDb(uid string, request *requests.CreateUserRequest) (*responses.User, error) {
	user := responses.User{
		Uid:       uid,
		Email:     request.Email,
		Name:      request.Name,
		Role:      request.Role,
		Disabled:  false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err := config.Firestore.Collection("users").Doc(uid).Set(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetAllUserFromDb() ([]*responses.User, error) {
	iter := config.Firestore.Collection("users").Documents(context.Background())
	documents, err := iter.GetAll()
	if err != nil {
		return nil, err
	}
	users := make([]*responses.User, 0)
	for _, doc := range documents {
		data := doc.Data()
		user := responses.UserFromJson(data)
		users = append(users, &user)
	}
	return users, nil
}

func DisableUserInDb(uid string) (*responses.User, error) {
	user, err := GetUserByUid(uid)
	if err != nil {
		return nil, err
	}
	user.Disabled = true
	_, err = config.Firestore.Collection("users").Doc(uid).Set(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByUid(uid string) (*responses.User, error) {
	doc, err := config.Firestore.Collection("users").Doc(uid).Get(context.Background())
	if err != nil {
		return nil, err
	}
	user := responses.UserFromJson(doc.Data())
	return &user, nil
}

func GetUserCount() (*int64, error) {
	query := config.Firestore.Collection("users").NewAggregationQuery().WithCount("count")
	data, err := query.Get(context.Background())
	if err != nil {
		return nil, err
	}
	value := data["count"].(*firestorepb.Value)
	count := value.GetIntegerValue()
	return &count, nil
}

func GetCasesCount() (*int64, error) {
	query := config.Firestore.Collection("cases").NewAggregationQuery().WithCount("count")
	data, err := query.Get(context.Background())
	if err != nil {
		return nil, err
	}
	value := data["count"].(*firestorepb.Value)
	count := value.GetIntegerValue()
	return &count, nil
}

func GetAllCases() ([]*responses.CaseModel, error) {
	iter := config.Firestore.Collection("cases").Documents(context.Background())
	documents, err := iter.GetAll()
	if err != nil {
		return nil, err
	}
	cases := make([]*responses.CaseModel, 0)
	for _, doc := range documents {
		data := doc.Data()
		caseModel := responses.CaseModelFromJson(data)
		hearings, err := getAllHearings(doc.Ref.ID)
		if err != nil {
			cases = append(cases, &caseModel)
		} else {
			caseModel.Hearings = append(caseModel.Hearings, *hearings...)
			cases = append(cases, &caseModel)
		}
	}
	return cases, nil
}

func getAllHearings(caseId string) (*[]responses.Hearing, error) {
	iter := config.Firestore.Collection("hearings").Where("case_id", "==", caseId).Documents(context.Background())
	documents, err := iter.GetAll()
	if err != nil {
		return nil, err
	}
	hearings := make([]responses.Hearing, 0)
	for _, doc := range documents {
		hearing := responses.HearingFromMap(doc.Data())
		hearings = append(hearings, hearing)
	}
	return &hearings, nil
}
