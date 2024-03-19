package handler

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/bjedrzejewsk/spotify-quiz/model"
	"github.com/bjedrzejewsk/spotify-quiz/service"
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
	parsedTracks := playlist.ParseFullTracks(rawTracks)
	playlist.SetTracks(parsedTracks)

	service.DisplaySongsTemplate(w, playlist.Tracks)
}

func GetPlaylistImage(w http.ResponseWriter, r *http.Request) {
	fullPlaylist, err := Client.GetPlaylist(spotify.ID(playlist.Id))
	if err != nil {
		log.Fatalln(err)
	}
	image := playlist.GetImageURL(*fullPlaylist)

	DisplayImageTemplete(w, image)
}

func DisplayImageTemplete(w http.ResponseWriter, image string) {
	htmlStr := fmt.Sprintf("<img src=%s></img>", image)
	tmpl, _ := template.New("t").Parse(htmlStr)
	tmpl.Execute(w, tmpl)
}
