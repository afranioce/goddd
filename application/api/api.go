package api

import (
	"github.com/afranioce/goddd/domain/eventsourcing"
	"github.com/afranioce/goddd/domain/eventsourcing/command"
	"github.com/afranioce/goddd/domain/eventsourcing/handler"
)

var dispatcher = eventsourcing.NewCommandRegister()

func init() {
	dispatcher.Add(command.CreateAccount{}, handler.CreateAccountHandler)
	dispatcher.Add(command.ChangePassword{}, handler.ChangePasswordHandler)
}
