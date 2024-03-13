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

func (playlist *Playlist) ParseTracks(rawTracks *spotify.PlaylistTrackPage) []Track {
	var tracks []Track
	for _, item := range rawTracks.Tracks {
		tracks = append(tracks, Track{
			Name:  item.Track.Name,
			Album: Album{ReleaseDate: item.Track.Album.ReleaseDate},
		})
	}
	return tracks
}

func (playlist *Playlist) SetTracks(tracks []Track) {
	playlist.Tracks = tracks
}

func (playlist *Playlist) GetImageURL(fullPlaylist spotify.FullPlaylist) string {
	firstImageIndex := 0
	mediumImageIndex := 1

	log.Print("Processing playlist image list: ", fullPlaylist.Images)

	if playlist.hasOneImage(fullPlaylist) {
		return fullPlaylist.Images[firstImageIndex].URL
	} else {
		return fullPlaylist.Images[mediumImageIndex].URL
	}
}

func (playlist *Playlist) hasOneImage(fullPlaylist spotify.FullPlaylist) bool {
	return len(fullPlaylist.Images) == 1
}
