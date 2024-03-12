package handler

import (
	"log"
	"math/rand"
	"net/http"

	"github.com/zmb3/spotify"
)

func RandomTrackPlayback(w http.ResponseWriter, r *http.Request) {
	log.Print("Playback toggled:")
	state, err := Client.PlayerState()
	if err != nil {
		log.Println(err)
	}
	if state.CurrentlyPlaying.Playing {
		TogglePlaybackState(state)
	}
	if breakIfNoActiveDevices() {
		return
	}

	uri := "spotify:playlist:" + spotify.URI(playlist.Id)
	offset := spotify.PlaybackOffset{Position: rand.Intn(len(playlist.Tracks))}
	playbackOptions := spotify.PlayOptions{PlaybackContext: &uri, PlaybackOffset: &offset}
	Client.PlayOpt(&playbackOptions)

}

func StartStopPlayback(w http.ResponseWriter, r *http.Request) {
	log.Print("Playback toggled:")
	state, err := Client.PlayerState()
	if err != nil {
		log.Println(err)
	}

	if breakIfNoActiveDevices() {
		return
	}

	TogglePlaybackState(state)
}

func breakIfNoActiveDevices() bool {
	devices, err := Client.PlayerDevices()
	if err != nil {
		log.Println(err)
	}

	if len(devices) == 0 {
		log.Print("No active user's devices found")
		return true
	}
	return false
}

func TogglePlaybackState(state *spotify.PlayerState) {
	switch state.CurrentlyPlaying.Playing {
	case true:
		log.Print("Stop playback")
		err := Client.Pause()
		if err != nil {
			log.Println(err)
		}
	case false:
		log.Print("Start playback")
		err := Client.Play()
		if err != nil {
			log.Println(err)
		}
	default:
		log.Print("Something went wrong: Playback state is neither playing nor paused.")
	}
}
