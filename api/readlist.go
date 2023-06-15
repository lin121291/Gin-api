package api

import (
	db "Gin-api/db/sqlc"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type appendlistRequect struct {
	UserID sql.NullInt32
	NewsID sql.NullInt32
}

func (server *Server) appendlistItem(ctx *gin.Context) {
	var req appendlistRequect

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.AppendListParams{
		UserID: req.UserID,
		NewsID: req.NewsID,
	}

	account, err := server.store.AppendList(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, account)
}

type deletelistRequect struct {
	id int32
}

func (server *Server) deletelistItem(ctx *gin.Context) {
	var req deletelistRequect

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = server.store.DeleteList(ctx, req.id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, nil)
}
