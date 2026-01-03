package config

import (
    "context"
    "log"
    firebase "firebase.google.com/go/v4"
    "google.golang.org/api/option"
)

func InitFirebase() *firebase.App {
    opt := option.WithServiceAccountFile("serviceAccountKey.json")
    app, err := firebase.NewApp(context.Background(), nil, opt)
    if err != nil {
        log.Fatalf("Firebase error: %v", err)
    }
    return app
}