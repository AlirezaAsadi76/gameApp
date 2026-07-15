package dto

type ProfileRequest struct {
	UserId uint `json:"user_id"`
}
type ProfileResponse struct {
	User UserInfo `json:"user"`
}
