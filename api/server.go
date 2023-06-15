package api

import (
	db "Gin-api/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	server.setupRouter()
	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/account", server.createAccount)
	router.DELETE("/account/:id", server.deleteAccount)

	router.POST("/news", server.createNews)
	router.GET("/news/top", server.topNews)
	//router.POST("news/:id", server.deleteNews)

	router.POST("/readlist", server.appendlistItem)
	router.DELETE("/readlist/:id", server.deletelistItem)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
