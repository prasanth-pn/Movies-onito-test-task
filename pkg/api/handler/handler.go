package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prasanth-pn/Movies-onito-test-task/pkg/domain"
	"github.com/prasanth-pn/Movies-onito-test-task/pkg/usecase/interfaces"
	"github.com/prasanth-pn/Movies-onito-test-task/pkg/utils"
)

type MoviesHandler struct {
	moviusecase interfaces.MovieusecaseInterface
}

func NewMovieHandler(movie interfaces.MovieusecaseInterface) *MoviesHandler {
	return &MoviesHandler{movie}
}
func (m *MoviesHandler) LongestDurationMovies(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pagenation := utils.Filter{
		Page: page,
		PageSize: 10,
	}
	movies,metadata, err := m.moviusecase.LongestDurationMovies(pagenation)
	if err != nil {
		c.IndentedJSON(401, gin.H{
			"error": "error happend while getting data",
		})
		return
	}
	c.IndentedJSON(200, gin.H{
		"metadata":metadata,
		"MovieList":movies,
	})

}
func (m *MoviesHandler)Addnewmovie(c *gin.Context){
	var movie domain.Movies

	if err:=c.BindJSON(&movie);err!=nil{
		c.IndentedJSON(401,gin.H{
			"error":"errot geting input",
		})
	}
	err:=m.moviusecase.Addnewmovie(movie)
	if err!=nil{
		c.IndentedJSON(401,gin.H{
			"Error":"movie is not added",
		})
		return
	}
	c.IndentedJSON(201,gin.H{
		"Message ":"SUCCESS",
		"DATA":movie,
	})
}

func (m *MoviesHandler)TopRatedMovies(c *gin.Context){
	page,_:=strconv.Atoi(c.Query("page"))
	pagenation:=utils.Filter{
		Page: page,
		PageSize: 10,
	}
	movies,metadata,err:=m.moviusecase.TopRatedMovies(pagenation)
	if err!=nil{
		c.IndentedJSON(401,gin.H{
			"Error":"cant't get the toprated movies oops",
		})
	}
c.IndentedJSON(200,gin.H{
	"message":"SUCCESS",
	"data":movies,
	"metadata":metadata,
})
}


func (m *MoviesHandler)MoviesWithSubTotal(c *gin.Context){
page ,_:=strconv.Atoi(c.Query("page"))
pagenation:=utils.Filter{
	Page: page,
	PageSize: 10,
}
movies,metadata,err:=m.moviusecase.GenreMoviesWithSubTotal(pagenation)
if err!=nil{
	c.JSON(401,gin.H{
		"Error":"error when fetching the movies",
	})
	return 
}

	c.IndentedJSON(200,gin.H{
		"message":"SUCCESS",
		"DATA":movies,
		"metadata":metadata,
	})
}
func (m *MoviesHandler)UpdateRunTimeMinutes(c *gin.Context){
	err:=m.moviusecase.UpdateRunTimeMinutes()
	if err!=nil{
		c.JSON(401,gin.H{
			"Error":"the data is not updated",
		})
		return
	}

c.JSON(200,gin.H{
	"message":"SUCCESS",
	"Message":"suceesfully updated the time",
})
}
