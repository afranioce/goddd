package command

type CreateAccount struct {
	Username      string `json:"username" binding:"required"`
	Email         string `json:"email" binding:"required"`
	PlainPassword string `json:"plain_password" binding:"required"`
}

type ChangePassword struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,nefield=OldPassword"`
}
