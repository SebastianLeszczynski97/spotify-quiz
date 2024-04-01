package model

import (
	"log"
	"math/rand"
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
	ImageUrl       string
}

type Playlist struct {
	Url            string
	Id             string
	Tracks         []TrackInfo
	ImageUrl       string
	TrackIndexes   []int
	LastTrackIndex int
}

func (playlist *Playlist) ParseId() {
	parsedPalylistUrl, err := url.Parse(playlist.Url)
	if err != nil {
		log.Println(err)
	}
	t := strings.Split(parsedPalylistUrl.Path, "/")
	playlist.Id = t[len(t)-1]
}

func (playlist *Playlist) SetUnusedIndexes() {
	indexes := make([]int, len(playlist.Tracks))
	for i := range indexes {
		indexes[i] = i
	}
	rand.Shuffle(len(indexes), func(i, j int) {
		indexes[i], indexes[j] = indexes[j], indexes[i]
	})
	playlist.TrackIndexes = indexes
}

func (playlist *Playlist) SetTracks(rawTracks []spotify.PlaylistTrack) {

	var tracks []TrackInfo
	for _, fullTrack := range rawTracks {
		track := TrackInfo{}
		track.SetTrackInfo(&fullTrack.Track)
		tracks = append(tracks, track)
	}
	playlist.Tracks = tracks
}

func (playlist *Playlist) GetPlaybackOptions() spotify.PlayOptions {
	uri := "spotify:playlist:" + spotify.URI(playlist.Id)
	offset := spotify.PlaybackOffset{Position: int(playlist.TrackIndexes[playlist.LastTrackIndex])}
	playbackOptions := spotify.PlayOptions{PlaybackContext: &uri, PlaybackOffset: &offset}
	playlist.IncrementTrackIndex()
	return playbackOptions
}

func (playlist *Playlist) IncrementTrackIndex() {
	if playlist.LastTrackIndex == (len(playlist.TrackIndexes) - 1) {
		playlist.LastTrackIndex = 0
	}
	playlist.LastTrackIndex++
}

func (playlist *Playlist) SetPlaylistImageURL(fullPlaylist *spotify.FullPlaylist) {
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

func (track *TrackInfo) SetTrackInfo(rawTrack *spotify.FullTrack) {
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
	track.ImageUrl = rawTrack.Album.Images[0].URL
}
