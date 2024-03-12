package handler

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/bjedrzejewsk/spotify-quiz/model"
	"github.com/zmb3/spotify"
)

var playlist model.Playlist

func init() {
	playlist.Url = "https://open.spotify.com/playlist/5rn1uqM3yaXf15HBAJEzs4"
	playlist.ParseId()
}

func SetPlaylist(w http.ResponseWriter, r *http.Request) {
	playlist.Url = r.PostFormValue("input-playlist")
	playlist.ParseId()
	log.Printf("Playlist added %s", playlist)
}

func GetPlaylistSongs(w http.ResponseWriter, r *http.Request) {
	rawTracks, err := Client.GetPlaylistTracks(spotify.ID(playlist.Id))
	if err != nil {
		log.Fatalln(err)
	}
	parsedTracks := playlist.ParseTracks(rawTracks)
	playlist.SetTracks(parsedTracks)

	DisplaySongsTemplete(w)
}

func DisplaySongsTemplete(w http.ResponseWriter) {
	for _, item := range playlist.Tracks {
		htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s %s</li>", item.Album.ReleaseDate, item.Name)
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, tmpl)
	}
}
