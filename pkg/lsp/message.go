package lsp

type Request struct {
	RPC    string `json:"jsonrpc"`
	ID     int    `json:"id"`
	Method string `json:"method"`

	// we will specify the types of the params in the requests later
	// Params
}
type Response struct {
	RPC string `json:"jsonrpc"`
	Id  *int   `json:"id,omitempty"`
	// result
	// error
}
type Notification struct {
	Method string `json:"method"`
	RPC    string `json:"jsonrpc"`
}
