package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	tokenService "github.com/bjedrzejewsk/spotify-quiz/pkg"
)

type Album struct {
	ReleaseDate string `json:"release_date"`
}

type Track struct {
	//there will be more fields here soon
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

	var playlist string = "2cHhJoYSQtwf20GkTEUJh4"

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
		playlist = r.PostFormValue("input-playlist")
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
		startPlaylistPlayback(playlist)
		log.Print("Started playback")
	}

	http.HandleFunc("/", initTemplate)
	http.HandleFunc("/set-playlist/", setPlaylistHandler)
	http.HandleFunc("/set-token/", setTokenHandler)
	http.HandleFunc("/get-playlist-songs/", getPlaylistSongsHandler)
	http.HandleFunc("/start-playback/", startPlaybackHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func startPlaylistPlayback(playlist string) []byte {
	client := &http.Client{}
	//temporarily
	endpointUrl := "https://api.spotify.com/v1/me/player/play"
	//values := map[string]string{"context_uri": fmt.Sprintf("spotify:playlist:%v", playlist), "offset": "{\"position\": 0}"}
	values := "{\"context_uri\": \"spotify:playlist:5ht7ItJgpBH7W6vJ5BqpPr\",\"offset\": {\"position\": 5},\"position_ms\": 0}"
	jsonValue, _ := json.Marshal(values)
	request, err := http.NewRequest("PUT", endpointUrl, bytes.NewBuffer(jsonValue))
	//request, err := http.NewRequest("PUT", endpointUrl, nil)
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Set("Content-Type", "application/json")
	responseBody := clientDoRequest(client, request)

	return responseBody
}

func getPlaylistSongs(playlist string) []Item {
	client := &http.Client{}
	//temporarily
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
