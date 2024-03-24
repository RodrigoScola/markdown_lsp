package lsp

type TextDocumentItem struct {
	TextDocumentIdentifier
	LanguageId string `json:"languageId"`
	Text       string `json:"text"`
}

type TextDocumentIdentifier struct {
	Uri string `json:"uri"`
}
type VersionDocumentIdentifier struct {
	TextDocumentIdentifier
	Version int `json:"version"`
}

type TextDocumentPositionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Position     Position               `json:"position"`
}

type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}
type Location struct {
	Uri   string `json:"uri"`
	Range Range  `json:"range"`
}
type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}