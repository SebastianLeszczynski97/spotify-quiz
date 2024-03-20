package handler

import (
	"log"
	"math/rand"
	"net/http"

	"github.com/bjedrzejewsk/spotify-quiz/model"
	"github.com/bjedrzejewsk/spotify-quiz/service"
	"github.com/zmb3/spotify"
)

func GetCurrentTrackInfo(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting currently playing song info:")
	if breakIfNoActiveDevices() {
		return
	}
	currentlyPlaying, err := Client.PlayerCurrentlyPlaying()
	if err != nil {
		log.Println(err)
		return
	}
	var track model.TrackInfo
	track.SetTrackInfo(currentlyPlaying.Item)
	log.Println(track)
	service.DisplayTrackInfoPanelTemplate(w, track)
}

func RandomTrackPlayback(w http.ResponseWriter, r *http.Request) {
	log.Println("Playback toggled:")
	state, err := Client.PlayerState()
	if err != nil {
		log.Println(err)
	}
	if state.CurrentlyPlaying.Playing {
		togglePlaybackState(state)
	}
	if breakIfNoActiveDevices() {
		return
	}

	uri := "spotify:playlist:" + spotify.URI(playlist.Id)
	offset := spotify.PlaybackOffset{Position: rand.Intn(len(playlist.Tracks))}
	playbackOptions := spotify.PlayOptions{PlaybackContext: &uri, PlaybackOffset: &offset}
	Client.PlayOpt(&playbackOptions)
	dummyTrack := model.TrackInfo{
		ImageUrl:    "https://th.bing.com/th/id/OIG1.lFSmScQIgKiQyHVQ0.8o?pid=ImgGn",
		Name:        "",
		ReleaseDate: "",
		AlbumName:   "",
	}

	service.DisplayTrackInfoPanelTemplate(w, dummyTrack)
}

func StartStopPlayback(w http.ResponseWriter, r *http.Request) {
	log.Println("Playback toggled:")
	state, err := Client.PlayerState()
	if err != nil {
		log.Println(err)
	}

	if breakIfNoActiveDevices() {
		return
	}

	togglePlaybackState(state)
}

func breakIfNoActiveDevices() bool {
	devices, err := Client.PlayerDevices()
	if err != nil {
		log.Println(err)
	}

	if len(devices) == 0 {
		log.Println("No user's devices found")
		return true
	}
	for _, device := range devices {
		if device.Active {
			log.Printf("Active device: %s", device.Name)
			return false
		}
	}
	log.Println("No active user's devices found")
	return true
}

func togglePlaybackState(state *spotify.PlayerState) {
	switch state.CurrentlyPlaying.Playing {
	case true:
		log.Println("Stop playback")
		err := Client.Pause()
		if err != nil {
			log.Println(err)
		}
	case false:
		log.Println("Start playback")
		err := Client.Play()
		if err != nil {
			log.Println(err)
		}
	default:
		log.Println("Something went wrong: Playback state is neither playing nor paused.")
	}
}
