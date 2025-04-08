package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/Nishivaly/go-reddit/v2/reddit"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

const tokenFile = "token.json"

var (
	clientID     string
	clientSecret string
	redirectURL  string
	oauthConfig  *oauth2.Config
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	clientID = os.Getenv("REDDIT_CLIENT_ID")
	clientSecret = os.Getenv("REDDIT_CLIENT_SECRET")
	redirectURL = os.Getenv("REDDIT_REDIRECT_URL")

	oauthConfig = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"*"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.reddit.com/api/v1/authorize",
			TokenURL: "https://www.reddit.com/api/v1/access_token",
		},
	}
}

func saveToken(token *oauth2.Token) error {
	file, err := os.Create(tokenFile)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(token)
}

func loadToken() (*oauth2.Token, error) {
	file, err := os.Open(tokenFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var token oauth2.Token
	if err := json.NewDecoder(file).Decode(&token); err != nil {
		return nil, err
	}
	return &token, nil
}

// Perform OAuth login flow (runs a web server and waits for Reddit redirect)
func loginAndGetToken() (*oauth2.Token, error) {
	state := "random-state" // Should be random in production

	// Open browser to Reddit login
	url := oauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline) + "&duration=permanent"
	fmt.Println("Opening browser for Reddit login...")

	var openCmd string
	switch runtime.GOOS {
	case "linux":
		openCmd = "xdg-open"
	case "windows":
		openCmd = "rundll32"
		url = fmt.Sprintf("url.dll,FileProtocolHandler %s", url)
	case "darwin":
		openCmd = "open"
	default:
		log.Printf("Unsupported OS: %s. Please open this URL manually: %s\n", runtime.GOOS, url)
		return nil, fmt.Errorf("unsupported OS")
	}
	exec.Command(openCmd, url).Start()

	// Wait for redirect
	codeCh := make(chan string)
	srv := &http.Server{Addr: ":8080"}

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("state") != state {
			http.Error(w, "State mismatch", http.StatusBadRequest)
			return
		}
		codeCh <- r.URL.Query().Get("code")
		w.Write([]byte("Login successful! You can close this tab."))
		go srv.Shutdown(context.Background())
	})

	go srv.ListenAndServe()
	code := <-codeCh

	// Exchange code for token
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	if err := saveToken(token); err != nil {
		log.Println("Warning: couldn't save token:", err)
	}

	return token, nil
}

// Returns a Reddit client using the token flow
func GetRedditClient() (*reddit.Client, error) {
	var token *oauth2.Token
	var err error

	token, err = loadToken()
	if err != nil {
		log.Println("No saved token found, starting login...")
		token, err = loginAndGetToken()
		if err != nil {
			return nil, fmt.Errorf("failed to get token: %w", err)
		}
	}

	// Build token source that will auto-refresh
	ts := oauthConfig.TokenSource(context.Background(), token)

	return reddit.NewClient(reddit.Credentials{}, reddit.WithTokenSource(ts))
}
