package handler

import (
	"log"
	"net/http"

	"github.com/bjedrzejewsk/spotify-quiz/model"
	"github.com/bjedrzejewsk/spotify-quiz/service"
)

type TemplateData struct {
	Tracks           []model.TrackInfo
	IsLogged         bool
	PlaylistImageUrl string
	IsPlaying        bool
}

func InitIndex(w http.ResponseWriter, r *http.Request) {
	Data := TemplateData{
		Tracks:           playlist.Tracks,
		IsLogged:         IsLoggedIn,
		PlaylistImageUrl: playlist.ImageUrl,
	}
	if IsLoggedIn {
		service.DisplayIndexPageTemplate(w, Data)
	} else {
		service.DisplayLoginPageTemplate(w, Data)
	}
}

func InitQuiz(w http.ResponseWriter, r *http.Request) {
	Data := TemplateData{
		Tracks:   playlist.Tracks,
		IsLogged: IsLoggedIn,
	}
	if IsLoggedIn {
		state, err := Client.PlayerState()
		if err != nil {
			log.Println(err)
		}
		Data.IsPlaying = state.CurrentlyPlaying.Playing
		service.DisplayQuizPageTemplate(w, Data)
	} else {
		service.DisplayLoginPageTemplate(w, Data)
	}
}
