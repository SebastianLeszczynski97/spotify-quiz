package handler

import (
	"fmt"
	"log"
	"net/http"
)

func StartStopPlayback(w http.ResponseWriter, r *http.Request) {
	log.Print("Playback toggled:")
	startStopPlayback()
}

func startStopPlayback() {
	//The client obj is declared and inicialised in authorization.go
	state, err := Client.PlayerState()
	if err != nil {
		log.Fatalln(err)
	}
	switch state.CurrentlyPlaying.Playing {
	case true:
		fmt.Println("Stop playback")
		err := Client.Pause()
		if err != nil {
			log.Fatalln(err)
		}
	case false:
		fmt.Println("Start playback")
		err := Client.Play()
		if err != nil {
			log.Fatalln(err)
		}
	default:
		fmt.Println("Something went wrong: Playback state is neither playing nor paused.")
	}
}
