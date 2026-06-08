package config

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"firebase.google.com/go/v4/messaging"
)

var Firestore *firestore.Client
var Auth *auth.Client
var MessagingClient *messaging.Client

func Init() error {
	config := &firebase.Config{
		ProjectID: "litigation-app-5f964",
	}
	app, err := firebase.NewApp(context.Background(), config)
	if err != nil {
		return err
	}
	client, err := app.Firestore(context.Background())
	if err != nil {
		return err
	}
	messagingClient, err := app.Messaging(context.Background())
	if err != nil {
		return err
	}
	authClient, err := app.Auth(context.Background())
	if err != nil {
		return err
	}
	Auth = authClient
	MessagingClient = messagingClient
	Firestore = client
	fmt.Println("Firestore created")
	return nil
}
