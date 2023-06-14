package api

import (
	db "Gin-api/db/sqlc"
	"fmt"
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
		fmt.Println("yoooooooooooo")
	}
	ctx.JSON(http.StatusOK, account)

}

func (server *Server) deleteAccount() {

}
