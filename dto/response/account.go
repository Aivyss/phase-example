package response

type PostAccountSignup struct {
	UserPK int `json:"id"`
}

func NewPostAccountSignup(userPK int) *PostAccountSignup {
	return &PostAccountSignup{
		UserPK: userPK,
	}
}
