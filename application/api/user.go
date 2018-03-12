package api

import (
	"net/http"

	"github.com/afranioce/goddd/domain/eventsourcing"
	"github.com/afranioce/goddd/domain/eventsourcing/command"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	Trans ut.Translator
}

func (c *User) Create(ctx *gin.Context) {
	account := command.CreateAccount{}

	if err := ctx.Bind(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, err.(validator.ValidationErrors).Translate(c.Trans))
		return
	}

	dispatcher.DispatchCommands(eventsourcing.CreateCommand(account, ctx))
}

func (c *User) ChangePassword(ctx *gin.Context) {
	changePassword := command.ChangePassword{}

	if err := ctx.Bind(&changePassword); err != nil {
		ctx.JSON(http.StatusBadRequest, err.(validator.ValidationErrors).Translate(c.Trans))
		return
	}

	dispatcher.DispatchCommands(eventsourcing.CreateCommand(changePassword, ctx))
}
