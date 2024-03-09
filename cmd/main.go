package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	tokenService "github.com/bjedrzejewsk/spotify-quiz/pkg"
	aut "github.com/bjedrzejewsk/spotify-quiz/pkg/authorization"
)

type Album struct {
	ReleaseDate string `json:"release_date"`
}
type Track struct {
	Name  string `json:"name"`
	Album Album  `json:"album"`
}
type Item struct {
	Track Track `json:"track"`
}
type PlaylistTrucksResponse struct {
	Items []Item `json:"items"`
}

func main() {
	fmt.Println("Go app... http://localhost:8080/")
	var playlistPlaceholder string = "https://open.spotify.com/playlist/5rn1uqM3yaXf15HBAJEzs4?si=de60d1492ec4484f"
	var playlist = parsePlaylsitId(playlistPlaceholder)

	initTemplate := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("../web/index.html"))
		tracks := map[string][]Track{
			"Tracks": {
				{Name: "Song name placeholder", Album: Album{ReleaseDate: "1997-05-22"}},
				{Name: "Song name placeholder2", Album: Album{ReleaseDate: "1997-11-11"}},
			},
		}
		tmpl.Execute(w, tracks)
	}

	setPlaylistHandler := func(w http.ResponseWriter, r *http.Request) {
		var rawPlaylist = r.PostFormValue("input-playlist")
		playlist = parsePlaylsitId(rawPlaylist)
		log.Printf(fmt.Sprintf("Playlist added %s", playlist), playlist)
	}

	setTokenHandler := func(w http.ResponseWriter, r *http.Request) {
		var newTokenValue = r.PostFormValue("input-token")
		tokenService.SetToken(newTokenValue)
		log.Printf(fmt.Sprintf("Token provided %s", tokenService.GetToken()), tokenService.GetToken())
	}

	getPlaylistSongsHandler := func(w http.ResponseWriter, r *http.Request) {
		playlistSongs := getPlaylistSongs(playlist)
		log.Print(playlistSongs)
		for _, item := range playlistSongs {
			htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s %s</li>", item.Track.Album.ReleaseDate, item.Track.Name)
			tmpl, _ := template.New("t").Parse(htmlStr)
			tmpl.Execute(w, tmpl)
		}
	}

	startPlaybackHandler := func(w http.ResponseWriter, r *http.Request) {
		startPlaylistPlayback()
		log.Print("Started playback")
	}

	http.HandleFunc("/", initTemplate)
	http.HandleFunc("/auth/login/", aut.Login)
	http.HandleFunc("/auth/callback/", aut.Callback)
	http.HandleFunc("/set-playlist/", setPlaylistHandler)
	http.HandleFunc("/set-token/", setTokenHandler)
	http.HandleFunc("/get-playlist-songs/", getPlaylistSongsHandler)
	http.HandleFunc("/start-playback/", startPlaybackHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func startPlaylistPlayback() {
	state, err := aut.Client.PlayerState()
	if err != nil {
		log.Fatalln(err)
	}
	switch state.CurrentlyPlaying.Playing {
	case true:
		fmt.Println("Stop playback")
		err := aut.Client.Pause()
		if err != nil {
			log.Fatalln(err)
		}
	default:
		fmt.Println("Start playback")
		err := aut.Client.Play()
		if err != nil {
			log.Fatalln(err)
		}
	}

}

func getPlaylistSongs(playlist string) []Item {
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

func parseTracks(body []byte) []Item {
	var unparsedResponse PlaylistTrucksResponse

	err := json.Unmarshal(body, &unparsedResponse)
	if err != nil {
		log.Fatalln("Error parsing JSON:", err)
	}
	return unparsedResponse.Items
}

func parsePlaylsitId(playlistLink string) string {
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
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", tokenService.GetToken()))
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
