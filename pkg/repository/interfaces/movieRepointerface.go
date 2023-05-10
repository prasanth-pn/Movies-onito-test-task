package interfaces

import (
	"github.com/prasanth-pn/Movies-onito-test-task/pkg/domain"
	"github.com/prasanth-pn/Movies-onito-test-task/pkg/utils"
)



type MoviesRepoInterfaces interface{
	LongestDurationMovies(pagenation utils.Filter)([]domain.MoviesResponse,utils.Metadata,error)
	Addnewmovie(movie domain.Movies)(error)
	TopRatedMovies(pagenation utils.Filter)([]domain.TopRateResponse,utils.Metadata,error)
	GenreMoviesWithSubTotal(pagenation utils.Filter)([]domain.SubtotalResponse,utils.Metadata,error)
	UpdateRunTimeMinutes()error
}