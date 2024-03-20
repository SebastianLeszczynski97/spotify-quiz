package handler

import (
	"net/http"

	"github.com/bjedrzejewsk/spotify-quiz/model"
	"github.com/bjedrzejewsk/spotify-quiz/service"
)

type TemplateData struct {
	Tracks           []model.TrackInfo
	IsLogged         bool
	PlaylistImageUrl string
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
		service.DisplayQuizPageTemplate(w, Data)
	} else {
		service.DisplayLoginPageTemplate(w, Data)
	}
}
