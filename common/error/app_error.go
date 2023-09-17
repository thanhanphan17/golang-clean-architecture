package cerr

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	RequestID  string `json:"request_id"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}
