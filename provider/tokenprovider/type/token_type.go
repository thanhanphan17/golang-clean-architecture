package tokentype

type TokenType int

const (
	VERIFY_TOKEN TokenType = iota
	ACCESS_TOKEN
	REFRESH_TOKEN
)

func (status TokenType) Value() string {
	switch status {
	case VERIFY_TOKEN:
		return "verify_token"
	case ACCESS_TOKEN:
		return "access_token"
	case REFRESH_TOKEN:
		return "refresh_token"
	}
	return ""
}
