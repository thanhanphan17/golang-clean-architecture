package requester

type JWTRequester struct {
	ID   string
	Role string
}

func (u JWTRequester) GetUserId() string {
	return u.ID
}

func (u JWTRequester) GetRole() string {
	return u.Role
}

var _ Requester = (*JWTRequester)(nil)
