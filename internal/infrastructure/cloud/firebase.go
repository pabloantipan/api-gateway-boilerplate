package cloud

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/pabloantipan/go-api-gateway-poc/config"
	"google.golang.org/api/option"
)

type FirebaseClient struct {
	auth      *auth.Client
	webAPIKey string
}

func NewFirebaseClient(ctx context.Context, cfg *config.Config) (*FirebaseClient, error) {
	opt := option.WithCredentialsFile(cfg.FirebaseCredentialsFile)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	return &FirebaseClient{
		auth:      authClient,
		webAPIKey: cfg.FirebaseWebAPIKey,
	}, nil
}

func (fc *FirebaseClient) VerifyToken(ctx context.Context, token string) (*auth.Token, error) {
	return fc.auth.VerifyIDToken(ctx, token)
}

func (fc *FirebaseClient) SignInWithPassword(ctx context.Context, email, password string) (string, error) {

	signInURL := fmt.Sprintf(
		"https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s",
		fc.webAPIKey,
	)

	reqBody := map[string]interface{}{
		"email":             email,
		"password":          password,
		"returnSecureToken": true,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", signInURL, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		IdToken string `json:"idToken"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.IdToken, nil
}
