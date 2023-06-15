package api

import (
	db "Gin-api/db/sqlc"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateNewsRequest struct {
	Title         string         `json:"title"`
	Content       sql.NullString `json:"content"`
	DatePublished sql.NullTime   `json:"date_published"`
}

func (server *Server) createNews(ctx *gin.Context) {
	var req CreateNewsRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateNewsParams{
		Title:         req.Title,
		Content:       req.Content,
		DatePublished: req.DatePublished,
	}

	account, err := server.store.CreateNews(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, account)
}

func (server *Server) topNews(ctx *gin.Context) {
	account, err := server.store.GetTopNews(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, account)
}
