package model

import (
	"log"
	"net/url"
	"strings"

	"github.com/zmb3/spotify"
)

type TrackInfo struct {
	Name           string
	ReleaseDate    string
	MainArtistName string
	ArtistNames    []string
	AlbumName      string
}

type Playlist struct {
	Url      string
	Id       string
	Tracks   []TrackInfo
	ImageUrl string
}

func (playlist *Playlist) ParseId() {
	parsedPalylistUrl, err := url.Parse(playlist.Url)
	if err != nil {
		log.Println(err)
	}
	t := strings.Split(parsedPalylistUrl.Path, "/")
	playlist.Id = t[len(t)-1]
}

func (playlist *Playlist) SetTracks(rawTracks *spotify.PlaylistTrackPage) {
	var tracks []TrackInfo
	for _, fullTrack := range rawTracks.Tracks {
		track := TrackInfo{}
		track.ParseFullTrack(&fullTrack.Track)
		tracks = append(tracks, track)
	}
	playlist.Tracks = tracks
}

func (playlist *Playlist) SetImageURL(fullPlaylist *spotify.FullPlaylist) {
	if hasNoImage(fullPlaylist) {
		log.Println("Playlist has no images")
		return
	}

	firstImageIndex := 0
	mediumImageIndex := 1
	log.Println("Processing playlist image list: ", fullPlaylist.Images)

	if hasOneImage(fullPlaylist) {
		playlist.ImageUrl = fullPlaylist.Images[firstImageIndex].URL
	} else {
		playlist.ImageUrl = fullPlaylist.Images[mediumImageIndex].URL
	}
}

func hasNoImage(fullPlaylist *spotify.FullPlaylist) bool {
	return len(fullPlaylist.Images) == 0
}

func hasOneImage(fullPlaylist *spotify.FullPlaylist) bool {
	return len(fullPlaylist.Images) == 1
}

func (track *TrackInfo) ParseFullTrack(rawTrack *spotify.FullTrack) {
	track.Name = rawTrack.Name
	track.ReleaseDate = rawTrack.Album.ReleaseDate
	var artists []string
	for index, astist := range rawTrack.Artists {
		if index == 0 {
			track.MainArtistName = astist.Name
		}
		artists = append(artists, astist.Name)
	}
	track.ArtistNames = artists
	track.AlbumName = rawTrack.Album.Name
}
