package entity

type UserStatus int

const (
	ACTIVE UserStatus = iota
	INACTIVE
	BLOCK
	VERIFIED
	NOT_VERIFIED
)

func (status UserStatus) Value() string {
	switch status {
	case ACTIVE:
		return "active"
	case INACTIVE:
		return "inactive"
	case BLOCK:
		return "block"
	case VERIFIED:
		return "verified"
	case NOT_VERIFIED:
		return "not_verified"
	}
	return ""
}
