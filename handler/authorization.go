package handler

import (
	"crypto/sha256"
	"encoding/base64"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bjedrzejewsk/spotify-quiz/service"
	"github.com/google/uuid"
	"github.com/zmb3/spotify"
)

func init() {
	auth = spotify.NewAuthenticator(
		os.Getenv("GO_SERVER_EXTERNAL_URL")+"/auth/callback",
		spotify.ScopeUserModifyPlaybackState,
		spotify.ScopeUserReadPlaybackState,
		spotify.ScopeUserReadPrivate)
	IsLoggedIn = false
}

var (
	auth       spotify.Authenticator
	Client     spotify.Client
	IsLoggedIn bool
)

func Login(w http.ResponseWriter, r *http.Request) {

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

	IsLoggedIn = true
	w.Header().Set("HX-Redirect", loginURL)
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
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	log.Printf("Logged as user: %s", user.DisplayName)
	IsLoggedIn = true
	http.Redirect(w, r, os.Getenv("GO_SERVER_EXTERNAL_URL"), http.StatusSeeOther)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	sessionID, err := r.Cookie("htssess")
	if err != nil {
		log.Printf("Error: %s", err.Error())
		IsLoggedIn = false
		http.Error(w, err.Error(), http.StatusInternalServerError)
		service.DisplayIndexPageTemplate(w, TemplateData{Tracks: playlist.Tracks, IsLogged: IsLoggedIn})
		return
	}

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
	log.Printf("Logged out")
	IsLoggedIn = false
	service.DisplayIndexPageTemplate(w, TemplateData{Tracks: playlist.Tracks, IsLogged: IsLoggedIn})
}

func getURL(state string) string {
	return auth.AuthURL(state)
}

func sha256Hash(input string) string {
	h := sha256.New()
	h.Write([]byte(input))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}
