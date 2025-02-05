package cloud

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"os"

// 	firebase "firebase.google.com/go/v4"
// 	"firebase.google.com/go/v4/auth"
// 	"google.golang.org/api/option"
// )

// type FirebaseWebClient struct {
// 	auth      *auth.Client
// 	webAPIKey string
// }

// func NewFirebaseWebClient(ctx context.Context, credentialsFile string) (*FirebaseWebClient, error) {
// 	opt := option.WithCredentialsFile(credentialsFile)
// 	app, err := firebase.NewApp(ctx, nil, opt)
// 	if err != nil {
// 		return nil, err
// 	}

// 	authClient, err := app.Auth(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Read web API key from environment or configuration
// 	webAPIKey := os.Getenv("FIREBASE_WEB_API_KEY")
// 	if webAPIKey == "" {
// 		return nil, fmt.Errorf("FIREBASE_WEB_API_KEY environment variable is required")
// 	}

// 	return &FirebaseWebClient{
// 		auth:      authClient,
// 		webAPIKey: "AIzaSyDxQPCvqDyenFVNPRi56P5pzvET09ryVMc", // webAPIKey,
// 	}, nil
// }

// func (fc *FirebaseWebClient) SignInWithPassword(ctx context.Context, email, password string) (string, error) {
// 	signInURL := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s", fc.webAPIKey)

// 	reqBody := map[string]interface{}{
// 		"email":             email,
// 		"password":          password,
// 		"returnSecureToken": true,
// 	}

// 	body, err := json.Marshal(reqBody)
// 	if err != nil {
// 		return "", err
// 	}

// 	req, err := http.NewRequestWithContext(ctx, "POST", signInURL, bytes.NewBuffer(body))
// 	if err != nil {
// 		return "", err
// 	}

// 	req.Header.Set("Content-Type", "application/json")

// 	resp, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	var result struct {
// 		IdToken string `json:"idToken"`
// 	}
// 	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
// 		return "", err
// 	}

// 	return result.IdToken, nil
// }
