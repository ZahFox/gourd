package command

// Result is what a client will recieve after issuing a command
type Response struct {
	ID    string      `json:"id"`
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}
