
package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	config util.Config
	router *gin.Engine
}

func NewServer(store *db.Store, config util.Config) (*Server, error) {
	server := &Server{store: store, config: config}

	server.setupRoutes()

	return server, nil
}

func (server *Server) setupRoutes() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"ping": "pong"})
	})

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}