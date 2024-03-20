package main

import (
	"log"
	"net/http"

	"github.com/bjedrzejewsk/spotify-quiz/handler"
)

func main() {
	log.Printf("Go app... http://localhost:8080/")

	http.HandleFunc("/", handler.InitIndex)
	http.HandleFunc("/quiz/", handler.InitQuiz)
	http.HandleFunc("/get-playlist-image/", handler.GetPlaylistImage)
	http.HandleFunc("/auth/login/", handler.Login)
	http.HandleFunc("/auth/logout/", handler.Logout)
	http.HandleFunc("/auth/callback/", handler.Callback)
	http.HandleFunc("/set-playlist/", handler.SetPlaylist)
	http.HandleFunc("/get-playlist-songs/", handler.GetPlaylistSongs)
	http.HandleFunc("/start-playback/", handler.StartStopPlayback)
	http.HandleFunc("/random-track-playback/", handler.RandomTrackPlayback)
	http.HandleFunc("/get-current-track-info/", handler.GetCurrentTrackInfo)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
