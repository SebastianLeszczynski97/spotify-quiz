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
	Url    string
	Id     string
	Tracks []TrackInfo
}

func (playlist *Playlist) ParseId() {
	parsedPalylistUrl, err := url.Parse(playlist.Url)
	if err != nil {
		log.Fatalln(err)
	}
	t := strings.Split(parsedPalylistUrl.Path, "/")
	playlist.Id = t[len(t)-1]
}

func (playlist *Playlist) ParseFullTracks(rawTracks *spotify.PlaylistTrackPage) []TrackInfo {
	var tracks []TrackInfo
	for _, fullTrack := range rawTracks.Tracks {
		track := TrackInfo{}
		track.ParseFullTrack(&fullTrack.Track)
		tracks = append(tracks, track)
	}

	return tracks
}

func (playlist *Playlist) SetTracks(tracks []TrackInfo) {
	playlist.Tracks = tracks
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

func (track *TrackInfo) ParseFullTrack1(rawTrack *spotify.FullTrack) {
	track.Name = rawTrack.Name
	track.ReleaseDate = rawTrack.Album.ReleaseDate
	var artists []string
	for index, astist := range rawTrack.Artists {
		if index == 1 {
			track.MainArtistName = astist.Name
		}
		artists = append(artists, astist.Name)
	}
	track.ArtistNames = artists
	track.AlbumName = rawTrack.Album.Name
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
