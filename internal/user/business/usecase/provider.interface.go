package usecase

type HashProvider interface {
	Hash(data string) string
}

type TokenProvider interface {
	Generate(payload map[string]interface{}, expiry uint) (map[string]interface{}, error)
}
