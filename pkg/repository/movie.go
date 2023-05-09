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
