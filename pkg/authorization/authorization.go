package authorization

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/zmb3/spotify"
)

func init() {
	auth = spotify.NewAuthenticator(
		os.Getenv("GO_SERVER_EXTERNAL_URL")+"/auth/callback",
		spotify.ScopeUserModifyPlaybackState,
		spotify.ScopeUserReadPlaybackState,
		spotify.ScopeUserReadPrivate)
}

var (
	auth   spotify.Authenticator
	Client spotify.Client
)

func Login(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, os.Getenv("GO_SERVER_EXTERNAL_URL")+"/dupa/", http.StatusSeeOther)
	sessionID := uuid.New().String()

	loginURL := getURL(sha256Hash(sessionID))

	expire := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{
		Name:    "htssess",
		Value:   sessionID,
		Path:    "/",
		Expires: expire,
	}
	if os.Getenv("HTTPS") == "true" {
		cookie.Secure = true
		cookie.SameSite = 4 //SameSite=None
	}

	http.SetCookie(w, &cookie)

	// w.Header().Set("HX-Redirect", loginURL)
	// w.WriteHeader(204)

	// http.Redirect(w, r, loginURL, http.StatusSeeOther)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", loginURL)
}

func getURL(state string) string {
	return auth.AuthURL(state)
}

func sha256Hash(input string) string {
	h := sha256.New()
	h.Write([]byte(input))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

func Callback(w http.ResponseWriter, r *http.Request) {
	sessionCookie, err := r.Cookie("htssess")
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	sessionID := sessionCookie.Value
	sessionHash := sha256Hash(sessionID)

	token, err := auth.Token(sessionHash, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	Client = auth.NewClient(token)
	user, err := Client.CurrentUser()
	fmt.Printf("Logged into a client with bday on: %s", user.Birthdate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	http.Redirect(w, r, os.Getenv("GO_SERVER_EXTERNAL_URL")+"dupa", http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	sessionID, err := r.Cookie("htssess")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Setting an expired cookie to remove from browsers
	expire := time.Now().Add(-24 * time.Hour)
	cookie := http.Cookie{
		Name:    "htssess",
		Value:   sessionID.Value,
		Path:    "/",
		Expires: expire,
	}
	if os.Getenv("HTTPS") == "true" {
		cookie.Secure = true
		cookie.SameSite = 4 //SameSite=None
	}

	http.SetCookie(w, &cookie)

	Client = spotify.NewClient(nil)
	http.Redirect(w, r, os.Getenv("GO_SERVER_EXTERNAL_URL"), http.StatusSeeOther)
}
