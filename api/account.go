package api

import (
	db "Gin-api/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUsersParams{
		Password: req.Password,
		Username: req.Username,
	}

	account, err := server.store.CreateUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	ctx.JSON(http.StatusOK, account)

}

type deleteAccountRequest struct {
	Userid int32
}

func (server *Server) deleteAccount(ctx *gin.Context) {
	var req deleteAccountRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.Idparams{
		Userid: req.Userid,
	}

	err = server.store.DeleteAccount(ctx, arg)

}
