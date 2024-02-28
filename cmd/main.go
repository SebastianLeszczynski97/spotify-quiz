package main

import (
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

	var apiToken string = "placeholder"
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
		playlistSongs := getPlaylistSongs(playlist, apiToken)
		log.Print(playlistSongs)
		for _, item := range playlistSongs {
			htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s %s</li>", item.Track.Album.ReleaseDate, item.Track.Name)
			tmpl, _ := template.New("t").Parse(htmlStr)
			tmpl.Execute(w, tmpl)
		}
	}

	http.HandleFunc("/", initTemplate)
	http.HandleFunc("/set-playlist/", setPlaylistHandler)
	http.HandleFunc("/set-token/", setTokenHandler)
	http.HandleFunc("/get-playlist-songs/", getPlaylistSongsHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func getPlaylistSongs(playlist string, token string) []Item {
	client := &http.Client{}
	//temporarily
	filter := "?fields=items%28track%28name%2C+album%28release_date%29%29%29"
	endpointUrl := fmt.Sprintf("https://api.spotify.com/v1/playlists/%s/tracks/%s", playlist, filter)
	request, err := http.NewRequest("GET", endpointUrl, nil)
	if err != nil {
		log.Fatalln(err)
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	fmt.Print(request.Header)
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	response := parseTracks(body)
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
