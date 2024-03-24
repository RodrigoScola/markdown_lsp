package lsp

type TextDocumentDidChangeNotification struct {
	Notification
	Params DidChangeTextDocumentNotificationParams `json:"params"`
}

type DidChangeTextDocumentNotificationParams struct {
	TextDocument   VersionDocumentIdentifier        `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

type TextDocumentContentChangeEvent struct {
	Text string `json:"text"`
}