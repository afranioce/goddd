package handler

import (
	"net/http"

	"github.com/afranioce/goddd/domain/entity"
	"github.com/afranioce/goddd/domain/eventsourcing"
	"github.com/afranioce/goddd/domain/eventsourcing/command"
	"github.com/afranioce/goddd/domain/eventsourcing/event"
	"github.com/afranioce/goddd/domain/eventsourcing/result"
	"github.com/afranioce/goddd/infraestructure/repository"
)

func CreateAccountHandler(cmd eventsourcing.Command) error {
	c, _ := cmd.Body().(command.CreateAccount)

	dom := entity.NewUser(c.Username, c.Email, c.PlainPassword)

	if err := dom.Check(); err != nil {
		cmd.Context().JSON(http.StatusBadRequest, err.Error())
		return nil
	}

	if err := repository.NewRepository().Save(dom); err != nil {
		return err
	}

	cmd.Context().JSON(http.StatusOK, result.User{
		ID:                  dom.ID,
		Email:               dom.Email,
		Username:            dom.Username,
		LastLogin:           *dom.LastLogin,
		Status:              dom.Status,
		PasswordRequestedAt: dom.PasswordRequestedAt,
	})

	return nil
}

func ChangePasswordHandler(cmd eventsourcing.Command) error {
	c, _ := cmd.Body().(command.ChangePassword)
	usr := &entity.User{}

	eventsourcing.LoadFromEvents(usr, []eventsourcing.Event{
		event.PasswordChanged{Password: c.NewPassword},
	})

	cmd.Context().JSON(http.StatusOK, nil)

	return nil
}
