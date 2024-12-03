package authDto

type ResetRequest struct {
	Email       string `json:"email"`
	OldPassword string `json:"old_password,omitempty"`
	NewPassword string `json:"new_password"`
}
