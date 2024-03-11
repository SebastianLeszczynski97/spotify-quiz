package model

import (
	"log"
	"net/url"
	"strings"

	"github.com/zmb3/spotify"
)

type Album struct {
	ReleaseDate string `json:"release_date"`
}
type Track struct {
	Name  string `json:"name"`
	Album Album  `json:"album"`
}
type Item struct {
	Track Track `json:"track"`
}
type PlaylistTrucksResponse struct {
	Items []Item `json:"items"`
}

type Playlist struct {
	Url    string
	Id     string
	Tracks []Track
}

func (playlist *Playlist) ParseId() {
	parsedPalylistUrl, err := url.Parse(playlist.Url)
	if err != nil {
		log.Fatalln(err)
	}
	t := strings.Split(parsedPalylistUrl.Path, "/")
	playlist.Id = t[len(t)-1]
}

func (playlist *Playlist) ParseTracks(rawTracks *spotify.PlaylistTrackPage) {
	var tracks []Track
	for _, item := range rawTracks.Tracks {
		tracks = append(tracks, Track{
			Name:  item.Track.Name,
			Album: Album{ReleaseDate: item.Track.Album.ReleaseDate},
		})
	}
	playlist.Tracks = tracks
}
