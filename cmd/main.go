package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/bjedrzejewsk/spotify-quiz/handler"
	"github.com/bjedrzejewsk/spotify-quiz/model"
)

func main() {
	log.Printf("Go app... http://localhost:8080/")

	initTemplate := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("../web/index.html"))
		tracks := map[string][]model.TrackInfo{
			"Tracks": {
				{Name: "Song name placeholder", ReleaseDate: "1997-05-22"},
				{Name: "Song name placeholder2", ReleaseDate: "1997-11-11"},
			},
		}
		tmpl.Execute(w, tracks)
	}

	http.HandleFunc("/", initTemplate)
	http.HandleFunc("/auth/login/", handler.Login)
	http.HandleFunc("/auth/callback/", handler.Callback)
	http.HandleFunc("/set-playlist/", handler.SetPlaylist)
	http.HandleFunc("/get-playlist-songs/", handler.GetPlaylistSongs)
	http.HandleFunc("/start-playback/", handler.StartStopPlayback)
	http.HandleFunc("/random-track-playback/", handler.RandomTrackPlayback)
	http.HandleFunc("/get-current-track-info/", handler.GetCurrentTrackInfo)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
