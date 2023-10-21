package entity

type UserStatus int

const (
	INACTIVE UserStatus = iota
	ACTIVE
	BLOCK
	VERIFIED
	UNVERIFIED
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
	case UNVERIFIED:
		return "unverified"
	}
	return ""
}
