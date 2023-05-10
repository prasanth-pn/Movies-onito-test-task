package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/prasanth-pn/Movies-onito-test-task/pkg/domain"
	"github.com/prasanth-pn/Movies-onito-test-task/pkg/repository/interfaces"
	"github.com/prasanth-pn/Movies-onito-test-task/pkg/utils"
)

type MovieDatabase struct {
	DB *sql.DB
}

func NewMovieRepository(db *sql.DB) interfaces.MoviesRepoInterfaces {
	return &MovieDatabase{DB: db}
}
func (c *MovieDatabase) LongestDurationMovies(pagenation utils.Filter) ([]domain.MoviesResponse, utils.Metadata, error) {
	var movies []domain.MoviesResponse
	query := `SELECT COUNT(*) OVER() ,tconst,primay_title,runtime_minutes,genres  FROM movies ORDER BY runtime_minutes DESC LIMIT $1 OFFSET $2;`

	rows, err := c.DB.Query(query, pagenation.Limit(), pagenation.Offset())
	if err != nil {
		log.Fatal(err)
		return nil, utils.Metadata{}, nil
	}
	var totalRecords int
	defer rows.Close()

	for rows.Next() {
		var movie domain.MoviesResponse
		err := rows.Scan(&totalRecords, &movie.Tconst, &movie.PrimaryTitle, &movie.RuntimeMinutes, &movie.Genres)
		if err != nil {
			return nil, utils.Metadata{}, err
		}
		movies = append(movies, movie)
	}
	fmt.Println(totalRecords)
	return movies, utils.ComputerMetadata(&totalRecords, &pagenation.Page, &pagenation.PageSize), nil

}
func (m *MovieDatabase) Addnewmovie(movie domain.Movies) error {
	query := `INSERT INTO movies (tconst,title_type,primay_title,runtime_minutes,genres)VALUES($1,$2,$3,$4,$5);`
	err := m.DB.QueryRow(query, movie.Tconst, movie.TitleType, movie.PrimayTitle, movie.RuntimeMinutes, movie.Genres).Err()
	return err
}

func (m *MovieDatabase) TopRatedMovies(pagenation utils.Filter) ([]domain.TopRateResponse, utils.Metadata, error) {
	var movies []domain.TopRateResponse
	query := `SELECT COUNT(*) OVER(),movies.tconst,movies.primay_title,ratings.average_rating FROM movies
	 INNER JOIN ratings ON movies.tconst=ratings.tconst WHERE average_rating>6 ORDER BY average_rating DESC LIMIT $1 OFFSET $2;`
	rows, err := m.DB.Query(query, pagenation.Limit(), pagenation.Offset())
	if err != nil {
		return nil, utils.Metadata{}, err
	}
	defer rows.Close()
	var totalRecords int
	for rows.Next() {
		var movie domain.TopRateResponse
		err := rows.Scan(&totalRecords, &movie.Tconst, &movie.PrimaryTitle, &movie.AverageRating)
		if err != nil {
			return nil, utils.Metadata{}, err
		}
		movies = append(movies, movie)
	}
	return movies, utils.ComputerMetadata(&totalRecords, &pagenation.Page, &pagenation.PageSize), nil

}
func (m *MovieDatabase) GenreMoviesWithSubTotal(pagenation utils.Filter) ([]domain.SubtotalResponse, utils.Metadata, error) {
	var movies []domain.SubtotalResponse
	query := `SELECT COUNT(*) OVER(),
    COALESCE(m.genres, 'Total') AS genres,
    COALESCE(m.primay_title, 'Total') AS primay_title,
    COALESCE(SUM(r.num_votes), SUM(num_votes)) AS num_votes
FROM movies m
LEFT JOIN ratings r ON m.tconst = r.tconst
GROUP BY GROUPING SETS ((m.genres, m.primay_title), (m.genres), ())
ORDER BY m.genres, m.primay_title NULLS LAST limit $1 offset $2;`
	rows, err := m.DB.Query(query, pagenation.Limit(), pagenation.Offset())
	if err != nil {
		return nil, utils.Metadata{}, err
	}

	defer rows.Close()
	var totalRecords int

	for rows.Next() {
		var movie domain.SubtotalResponse
		err := rows.Scan(&totalRecords, &movie.Genres, &movie.PrimaryTitle, &movie.NumVotes)

		if err != nil {
			return nil, utils.Metadata{}, err

		}
		movies = append(movies, movie)
	}
	return movies, utils.ComputerMetadata(&totalRecords, &pagenation.Page, &pagenation.PageSize), nil

}
func (m *MovieDatabase) UpdateRunTimeMinutes() error {
	query := `UPDATE movies 
	SET runtime_minutes = 
	  CASE 
		WHEN genres = 'Documentary' THEN runtime_minutes + 15
		WHEN genres = 'Animation' THEN runtime_minutes + 30
		ELSE runtime_minutes+45
	  END;`
	err := m.DB.QueryRow(query).Err()
	fmt.Println(err)
	return err
}
