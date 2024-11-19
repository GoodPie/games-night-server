package firebase

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func InitFirebase(credPath string) (*firebase.App, error) {
   opt := option.WithCredentialsFile(credPath)
   return firebase.NewApp(context.Background(), nil, opt)
}