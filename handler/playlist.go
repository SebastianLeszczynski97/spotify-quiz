package handler

import (
	"log"
	"net/http"

	"github.com/bjedrzejewsk/spotify-quiz/model"
	"github.com/bjedrzejewsk/spotify-quiz/service"
	"github.com/zmb3/spotify"
)

var playlist model.Playlist

func init() {
	playlist.Url = "https://open.spotify.com/playlist/5rn1uqM3yaXf15HBAJEzs4"
	playlist.ParseId()
	playlist.Tracks = []model.TrackInfo{
		{Name: "Song name placeholder", ReleaseDate: "1997-05-22"},
		{Name: "Song name placeholder2", ReleaseDate: "1997-11-11"},
	}
}

func SetPlaylist(w http.ResponseWriter, r *http.Request) {
	playlist.Url = r.PostFormValue("input-playlist")
	playlist.ParseId()

	rawTracks, err := Client.GetPlaylistTracks(spotify.ID(playlist.Id))
	if err != nil {
		log.Println(err)
	}
	playlist.SetTracks(rawTracks)

	fullPlaylist, err := Client.GetPlaylist(spotify.ID(playlist.Id))
	if err != nil {
		log.Println(err)
	}
	playlist.SetImageURL(fullPlaylist)

	log.Printf("Playlist added %s", playlist)
}

func GetPlaylistSongs(w http.ResponseWriter, r *http.Request) {
	if len(playlist.Tracks) == 0 {
		log.Println("There are no tracks in the selected playlist")
	}

	service.DisplaySongsTemplate(w, playlist.Tracks)
}

func GetPlaylistImage(w http.ResponseWriter, r *http.Request) {
	if playlist.ImageUrl == "" {
		log.Println("Selected playlist has no images")
	}

	service.DisplayImageTemplate(w, playlist.ImageUrl)
}
