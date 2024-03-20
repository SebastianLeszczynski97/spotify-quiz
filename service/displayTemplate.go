package service

import (
	"fmt"
	"html/template"
	"net/http"
	"reflect"

	"github.com/bjedrzejewsk/spotify-quiz/model"
)

func DisplayTrackInfoTemplate(w http.ResponseWriter, trackInfo model.TrackInfo) {
	val := reflect.ValueOf(trackInfo)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	typ := val.Type()

	for i := 0; i < typ.NumField()-1; i++ {
		field := typ.Field(i)
		value := val.Field(i)
		htmlStr := fmt.Sprintf("<li class='list-group-item bg-secondary text-white'>%s: %s</li>", field.Name, value)
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, tmpl)
	}
}

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

func DisplayAlbumImageTemplate(w http.ResponseWriter, image string) {
	htmlStr := fmt.Sprintf("<img src=%s class='img-fluid rounded-3 shadow' alt='album or placeholder image'></img>", image)
	tmpl, _ := template.New("t").Parse(htmlStr)
	tmpl.Execute(w, tmpl)
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
