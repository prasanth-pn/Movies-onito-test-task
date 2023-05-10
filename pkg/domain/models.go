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

type MoviesResponse struct {
	Tconst         string
	PrimaryTitle   string
	RuntimeMinutes int
	Genres         string
}

type TopRateResponse struct {
	Tconst        string
	PrimaryTitle  string
	Genres        string
	AverageRating float32
}
type SubtotalResponse struct{
	Genres string 
	PrimaryTitle string
	NumVotes int
}

SELECT  tconst,primay_title,runtime_minutes,genres  FROM movies ORDER BY runtime_minutes DESC LIMIT 10 OFFSET 0;



