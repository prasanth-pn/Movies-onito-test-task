package interfaces

import (
	"github.com/prasanth-pn/Movies-onito-test-task/pkg/domain"
	"github.com/prasanth-pn/Movies-onito-test-task/pkg/utils"
)

type MovieusecaseInterface interface {
	LongestDurationMovies(pagenation utils.Filter) ([]domain.MoviesResponse,utils.Metadata, error)
}
