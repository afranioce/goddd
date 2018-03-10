package result

import (
	"time"

	"github.com/afranioce/goddd/domain/entity"
)

type User struct {
	ID                  uint          `json:"id"`
	Username            string        `json:"username"`
	Email               string        `json:"email"`
	LastLogin           time.Time     `json:"last_login"`
	PasswordRequestedAt time.Time     `json:"password_requested_at"`
	Status              entity.Status `json:"status"`
}
