package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/bjedrzejewsk/spotify-quiz/model"
	obsoleteTokenService "github.com/bjedrzejewsk/spotify-quiz/pkg"
)

var playlist string

func init() {
	var playlistPlaceholder string = "https://open.spotify.com/playlist/5rn1uqM3yaXf15HBAJEzs4"
	playlist = ParsePlaylistId(playlistPlaceholder)
}

func SetPlaylist(w http.ResponseWriter, r *http.Request) {
	var rawPlaylist = r.PostFormValue("input-playlist")
	playlist = ParsePlaylistId(rawPlaylist)
	log.Printf(fmt.Sprintf("Playlist added %s", playlist), playlist)
}

func GetPlaylistSongs(w http.ResponseWriter, r *http.Request) {
	playlistSongs := getPlaylistSongs(playlist)
	log.Print(playlistSongs)
	for _, item := range playlistSongs {
		htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s %s</li>", item.Track.Album.ReleaseDate, item.Track.Name)
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, tmpl)
	}
}

// Turn into reciver func for dedicated playlist model, then move to model file or new one
func ParsePlaylistId(playlistLink string) string {
	playlistUrl, err := url.Parse(playlistLink)
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Println(playlistUrl.Path)
	//fmt.Println(strings.Split(playlistUrl.Path, "/"))
	t := strings.Split(playlistUrl.Path, "/")
	var playlistId = t[len(t)-1]
	return playlistId
}

func clientDoRequest(client *http.Client, request *http.Request) []byte {
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", obsoleteTokenService.GetToken()))
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return responseBody
}

func getPlaylistSongs(playlist string) []model.Item {
	client := &http.Client{}
	filter := "?fields=items%28track%28name%2C+album%28release_date%29%29%29"
	endpointUrl := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s/tracks/%s", playlist, filter)
	request, err := http.NewRequest("GET", endpointUrl, nil)
	if err != nil {
		log.Fatalln(err)
	}
	responseBody := clientDoRequest(client, request)
	response := parseTracks(responseBody)
	return response
}

func parseTracks(body []byte) []model.Item {
	var unparsedResponse model.PlaylistTrucksResponse

	err := json.Unmarshal(body, &unparsedResponse)
	if err != nil {
		log.Fatalln("Error parsing JSON:", err)
	}
	return unparsedResponse.Items
}
