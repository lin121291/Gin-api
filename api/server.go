package api

import "github.com/gin-gonic/gin"

type Server struct {
	store  db.Store
	router *gin.Engine
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/user/re")

	server.router = router
}

func NewServer() *Server {
	server := &Server{}
	server.setupRouter()
	return server
}
