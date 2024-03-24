package lsp

type InitializeRequest struct {
	Request
	Params InitializaRequestParams `json:"params"`
}

type InitializaRequestParams struct {
	ClientInfo *ClientInfo `json:"clientInfo"`
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeResponse struct {
	Response
	Result InitializeResult `json:"result"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverInfo"`
}

type CompletionProvider struct {
	TriggerCharacters []string `json:"triggerCharacters"`
}

type ServerCapabilities struct {
	TextDocumnetSyncKind int                `json:"textDocumentSync"`
	HoverProvider        bool               `json:"hoverProvider"`
	DefinitionProvider   bool               `json:"definitionProvider"`
	CompletionProvider   CompletionProvider `json:"completionProvider"`
}

type TextDocumnetSyncKind struct{}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewInitializeResponse(id int) *InitializeResponse {
	return &InitializeResponse{
		Response: Response{
			RPC: "2.0",
			Id:  &id,
		},
		Result: InitializeResult{
			Capabilities: ServerCapabilities{
				TextDocumnetSyncKind: 1,
				HoverProvider:        true,
				DefinitionProvider:   true,
				CompletionProvider: CompletionProvider{
					TriggerCharacters: []string{"#", " ", ".", "<", ">", "/", "*", "+", "-", "&", "|", "!", "=", "%", "^", "$", "@", ":", ";", ",", "?", "[", "]", "{", "}", "(", ")", "\"", "'", "`", "~"},
				},
			},
			ServerInfo: ServerInfo{
				Name:    "mylsp",
				Version: "0.0.0.0.0.0-beta1.final",
			},
		},
	}

}