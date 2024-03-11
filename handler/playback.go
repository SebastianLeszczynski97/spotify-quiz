package handler

import (
	"log"
	"net/http"

	"github.com/zmb3/spotify"
)

func StartStopPlayback(w http.ResponseWriter, r *http.Request) {
	log.Print("Playback toggled:")
	state, err := Client.PlayerState()
	if err != nil {
		log.Fatalln(err)
	}

	devices, err := Client.PlayerDevices()
	if err != nil {
		log.Fatalln(err)
	}
	if len(devices) != 0 {
		TogglePlaybackState(state)
	} else {
		log.Print("No active user's devices found")
	}
}

func TogglePlaybackState(state *spotify.PlayerState) {
	switch state.CurrentlyPlaying.Playing {
	case true:
		log.Print("Stop playback")
		err := Client.Pause()
		if err != nil {
			log.Fatalln(err)
		}
	case false:
		log.Print("Start playback")
		err := Client.Play()
		if err != nil {
			log.Fatalln(err)
		}
	default:
		log.Print("Something went wrong: Playback state is neither playing nor paused.")
	}
}
