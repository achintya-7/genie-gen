package processes

import "os"

func Api() error {
	err := os.Mkdir("api", 0755)
	if err != nil {
		return err
	}

	err = os.Chdir("api")
	if err != nil {
		return err
	}

	file, err := os.Create("api.go")
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(`
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
}`)

	err = os.Chdir("..")
	if err != nil {
		return err
	}

	return nil
}
