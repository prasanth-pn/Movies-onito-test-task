package http

import (
	"github.com/gin-gonic/gin"
	"github.com/prasanth-pn/Movies-onito-test-task/pkg/api/handler"
	"github.com/prasanth-pn/Movies-onito-test-task/pkg/config"
)

type ServerHttp struct {
	engine *gin.Engine
}

func NewServerHttp(movieHandler *handler.MoviesHandler) *ServerHttp {
	engine := gin.New()
	//engine.Use(gin.Logger())
	api := engine.Group("api/v1")
	api.GET("/longest-duration-movies", movieHandler.LongestDurationMovies)
	api.POST("/new-movie",movieHandler.Addnewmovie)
	// api.GET("/top-rated-movies")
	// api.GET("/genre-movies-with-subtotals")
	// api.POST("/update-runtime-minutes")
	return &ServerHttp{engine: engine}
}
func (s *ServerHttp) Start(cf config.Config) {
	s.engine.Run(cf.Port)
}
