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
	playlist.ImageUrl = ""
	playlist.LastTrackIndex = 0
}

func SetPlaylist(w http.ResponseWriter, r *http.Request) {
	playlist.Url = r.PostFormValue("input-playlist")
	playlist.ParseId()
	response, err := Client.GetPlaylistTracks(spotify.ID(playlist.Id))
	if err != nil {
		log.Println(err)
	}
	var rawTracks []spotify.PlaylistTrack
	if response.Total > 0 {
		for i := 0; i < response.Total; i = i + 100 {
			responceTracks, err := Client.GetPlaylistTracksOpt(spotify.ID(playlist.Id), &spotify.Options{Offset: &i}, "")
			if err != nil {
				log.Println(err)
			}
			rawTracks = append(rawTracks, responceTracks.Tracks...)
		}
	}
	playlist.SetTracks(rawTracks)
	playlist.SetUnusedIndexes()
	fullPlaylist, err := Client.GetPlaylist(spotify.ID(playlist.Id))
	if err != nil {
		log.Println(err)
	}
	playlist.SetPlaylistImageURL(fullPlaylist)

	log.Printf("Playlist added %v", playlist)
	service.DisplayPlaylistImageTemplate(w, playlist.ImageUrl)
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

	service.DisplayPlaylistImageTemplate(w, playlist.ImageUrl)
}
