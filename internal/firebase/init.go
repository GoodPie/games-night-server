package firebase

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func InitFirebase(credPath string) (*firebase.App, error) {
	opt := option.WithCredentialsFile(credPath)
	return firebase.NewApp(context.Background(), nil, opt)
}

func GetAuthClient(app *firebase.App) (*auth.Client, error) {
	return app.Auth(context.Background())
}
