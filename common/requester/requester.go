package requester

const (
	CurrentRequester = ""
)

type Requester interface {
	GetUserId() string
	GetRole() string
}
