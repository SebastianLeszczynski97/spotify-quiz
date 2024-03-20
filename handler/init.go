package handler

import (
	"net/http"

	"github.com/bjedrzejewsk/spotify-quiz/model"
	"github.com/bjedrzejewsk/spotify-quiz/service"
)

type TemplateData struct {
	Tracks   []model.TrackInfo
	IsLogged bool
}

func InitIndex(w http.ResponseWriter, r *http.Request) {
	Data := TemplateData{
		Tracks:   playlist.Tracks,
		IsLogged: IsLoggedIn,
	}
	if IsLoggedIn {
		service.DisplayIndexPageTemplate(w, Data)
	} else {
		service.DisplayLoginPageTemplate(w, Data)
	}
}
