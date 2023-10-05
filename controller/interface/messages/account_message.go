package messages

import "github.com/xkurozaru/plant-diagnosis/controller/domain/model"

////////////////////////
// Request & Response //
////////////////////////

type SignUpRequest struct {
	Name     string `json:"name"`
	LoginID  string `json:"login_id"`
	Password string `json:"password"`
}
type SignUpResponse struct{}

type SignInRequest struct {
	LoginID  string `json:"login_id"`
	Password string `json:"password"`
}
type SignInResponse struct {
	Token string `json:"token"`
}

type GetUserRequest struct{}
type GetUserResponse struct {
	User UserMessage `json:"user"`
}

///////////
// Model //
///////////

type UserMessage struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	LoginID string `json:"login_id"`
}

func NewUserMessage(u model.User) UserMessage {
	return UserMessage{
		ID:      u.ID.ToString(),
		Name:    u.Name,
		LoginID: u.LoginID,
	}
}
