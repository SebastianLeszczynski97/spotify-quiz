package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

type CatFact struct {
	Fact string
}

type UnprocessedCatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func main() {
	fmt.Println("Go app...")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("../web/index.html"))
		catFacts := map[string][]CatFact{
			"CatFacts": {
				{Fact: "Cats are cool."},
				{Fact: "Cats eat mice."},
			},
		}
		tmpl.Execute(w, catFacts)
	}

	h2 := func(w http.ResponseWriter, r *http.Request) {
		playlist := r.PostFormValue("playlist")
		log.Printf(fmt.Sprintf("Playlist added %s", playlist))
	}

	h3 := func(w http.ResponseWriter, r *http.Request) {
		catFact := getCatFact()
		log.Printf(catFact)

		htmlStr := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s</li>", catFact)
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, tmpl)
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/add-playlist/", h2)
	http.HandleFunc("/show-cat-fact/", h3)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func getCatFact() string {
	resp, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	catFact := parseCatFact(body)
	return catFact
}

func parseCatFact(body []byte) string {
	var unpCatFact UnprocessedCatFact

	err := json.Unmarshal(body, &unpCatFact)
	if err != nil {
		log.Fatalln("Error parsing JSON:", err)
	}
	return unpCatFact.Fact
}
