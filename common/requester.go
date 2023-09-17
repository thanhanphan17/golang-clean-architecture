package common

const (
	CurrentRequester = ""
)

type Requester interface {
	GetUserId() string
	GetRole() string
}

var _ Requester = (*JWTRequesterData)(nil)

type JWTRequesterData struct {
	ID   string
	Role string
}

func (u JWTRequesterData) GetUserId() string {
	return u.ID
}

func (u JWTRequesterData) GetRole() string {
	return u.Role
}
