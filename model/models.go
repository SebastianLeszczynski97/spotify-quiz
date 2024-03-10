package model

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
