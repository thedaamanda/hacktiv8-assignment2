package responses

type Error struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type ErrorValidation struct {
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}
