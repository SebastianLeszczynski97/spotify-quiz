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
	log.Print("Getting currently playing song info:")
	if breakIfNoActiveDevices() {
		return
	}
	currentlyPlaying, err := Client.PlayerCurrentlyPlaying()
	if err != nil {
		log.Println(err)
		return
	}
	var track model.TrackInfo
	track.ParseFullTrack(currentlyPlaying.Item)
	log.Print(track)
	service.DisplayTrackInfoTemplate(w, track)
}

func RandomTrackPlayback(w http.ResponseWriter, r *http.Request) {
	log.Print("Playback toggled:")
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

	togglePlaybackState(state)
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

func togglePlaybackState(state *spotify.PlayerState) {
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
