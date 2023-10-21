package common

type AppRole int

const (
	ADMIN AppRole = iota
	LECTURER
	STUDENT
)

func (role AppRole) Value() string {
	switch role {
	case ADMIN:
		return "admin"
	case LECTURER:
		return "lecturer"
	case STUDENT:
		return "student"
	}
	return ""
}
