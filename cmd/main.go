package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/bjedrzejewsk/spotify-quiz/handler"
	"github.com/bjedrzejewsk/spotify-quiz/model"
)

func main() {
	fmt.Println("Go app... http://localhost:8080/")

	initTemplate := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("../web/index.html"))
		tracks := map[string][]model.Track{
			"Tracks": {
				{Name: "Song name placeholder", Album: model.Album{ReleaseDate: "1997-05-22"}},
				{Name: "Song name placeholder2", Album: model.Album{ReleaseDate: "1997-11-11"}},
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

	log.Fatal(http.ListenAndServe(":8080", nil))

}
