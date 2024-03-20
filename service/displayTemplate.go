package service

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/bjedrzejewsk/spotify-quiz/model"
)

func DisplaySongsTemplate(w http.ResponseWriter, tracks []model.TrackInfo) {
	for _, item := range tracks {
		htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s %s</li>", item.ReleaseDate, item.Name)
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, tmpl)
	}
}

func DisplayPlaylistImageTemplate(w http.ResponseWriter, image string) {
	htmlStr := fmt.Sprintf("<img src=%s class='img-fluid rounded-3 shadow' alt='album or placeholder image'></img>", image)
	tmpl, _ := template.New("t").Parse(htmlStr)
	tmpl.Execute(w, tmpl)
}

func DisplayTrackInfoPanelTemplate(w http.ResponseWriter, Track model.TrackInfo) {
	tmpl := template.Must(template.ParseFiles("../web/trackInfoPanel.html"))

	tmpl.Execute(w, Track)
}

func DisplayLoginPageTemplate(w http.ResponseWriter, data any) {
	tmpl := make(map[string]*template.Template)
	tmpl["index.html"] = template.Must(template.ParseFiles("../web/index.html", "../web/base.html"))
	tmpl["login.html"] = template.Must(template.ParseFiles("../web/login.html", "../web/base.html"))
	tmpl["quiz.html"] = template.Must(template.ParseFiles("../web/quiz.html", "../web/base.html"))

	tmpl["login.html"].ExecuteTemplate(w, "base", data)
}

func DisplayIndexPageTemplate(w http.ResponseWriter, data any) {
	tmpl := make(map[string]*template.Template)
	tmpl["index.html"] = template.Must(template.ParseFiles("../web/index.html", "../web/base.html"))
	tmpl["login.html"] = template.Must(template.ParseFiles("../web/login.html", "../web/base.html"))
	tmpl["quiz.html"] = template.Must(template.ParseFiles("../web/quiz.html", "../web/base.html"))

	tmpl["index.html"].ExecuteTemplate(w, "base", data)
}

func DisplayQuizPageTemplate(w http.ResponseWriter, data any) {
	tmpl := make(map[string]*template.Template)
	tmpl["index.html"] = template.Must(template.ParseFiles("../web/index.html", "../web/base.html"))
	tmpl["login.html"] = template.Must(template.ParseFiles("../web/login.html", "../web/base.html"))
	tmpl["quiz.html"] = template.Must(template.ParseFiles("../web/quiz.html", "../web/base.html"))

	tmpl["quiz.html"].ExecuteTemplate(w, "base", data)
}
