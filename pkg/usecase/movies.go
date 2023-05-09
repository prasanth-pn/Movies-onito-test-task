package usecase

import (
	"github.com/prasanth-pn/Movies-onito-test-task/pkg/domain"
	"github.com/prasanth-pn/Movies-onito-test-task/pkg/repository/interfaces"
	U "github.com/prasanth-pn/Movies-onito-test-task/pkg/usecase/interfaces"
	"github.com/prasanth-pn/Movies-onito-test-task/pkg/utils"
)

type MovieUseCase struct{
	Repo interfaces.MoviesRepoInterfaces
}
func NewMovieUseCase(repo interfaces.MoviesRepoInterfaces)U.MovieusecaseInterface{
	return &MovieUseCase{repo}
}
func (m *MovieUseCase)LongestDurationMovies(pagenation utils.Filter)([]domain.MoviesResponse,utils.Metadata,error){
	movies,metadata,err:=m.Repo.LongestDurationMovies(pagenation)
	return movies,metadata,err
}