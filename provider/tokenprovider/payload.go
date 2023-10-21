package tokenprovider

type TokenPayload struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	Type   string `json:"type"`
}
