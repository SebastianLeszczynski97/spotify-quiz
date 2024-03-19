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

	for i := 0; i < typ.NumField(); i++ {
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

func DisplayImageTemplete(w http.ResponseWriter, image string) {
	htmlStr := fmt.Sprintf("<img src=%s></img>", image)
	tmpl, _ := template.New("t").Parse(htmlStr)
	tmpl.Execute(w, tmpl)
}
