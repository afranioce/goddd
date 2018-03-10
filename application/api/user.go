package api

import (
	"net/http"

	"github.com/afranioce/goddd/domain/eventsourcing"
	"github.com/afranioce/goddd/domain/eventsourcing/command"
	"github.com/gin-gonic/gin"
)

type User struct{}

func (c *User) Create(ctx *gin.Context) {
	account := command.CreateAccount{}

	if err := ctx.Bind(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	dispatcher.DispatchCommands(eventsourcing.CreateCommand(account, ctx))
}

func (c *User) ChangePassword(ctx *gin.Context) {
	changePassword := command.ChangePassword{}

	if err := ctx.Bind(&changePassword); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	dispatcher.DispatchCommands(eventsourcing.CreateCommand(changePassword, ctx))
}
