package domain

type Movies struct {
	Tconst         string `json:"id"`
	TitleType      string `json:"title_type"`
	PrimayTitle    string `json:"primary_title"`
	RuntimeMinutes int    `json:"run_time_min"`
	Genres         string `json:"genres"`
}
type Ratings struct {
	Tconst        string  `json:"id"`
	AverageRating float32 `json:"average_rating"`
	NumVotes      int     `json:"num_votes"`
}


type MoviesResponse struct{
	Tconst string
	PrimaryTitle string
	RuntimeMinutes int
	Genres string
}
