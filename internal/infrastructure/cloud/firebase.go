package cloud

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

type FirebaseClient struct {
	auth *auth.Client
}

func NewFirebaseClient(ctx context.Context, credentialsFile string) (*FirebaseClient, error) {
	opt := option.WithCredentialsFile(credentialsFile)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	return &FirebaseClient{
		auth: authClient,
	}, nil
}

func (fc *FirebaseClient) VerifyToken(ctx context.Context, token string) (*auth.Token, error) {
	return fc.auth.VerifyIDToken(ctx, token)
}
