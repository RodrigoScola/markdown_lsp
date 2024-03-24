package analysis

import (
	"fmt"
	"log"
	"mylsp/pkg/lsp"
	"strings"
)

type State struct {
	// map of filenames to contents
	Documents map[string]string
}

func NewState() State {
	return State{Documents: map[string]string{}}
}

func getDiagnosticsForFileText(text string) []lsp.Diagnostics{

	diagnostics := []lsp.Diagnostics{}

	for row, line := range strings.Split(text, "\n") {
		


		if strings.Contains(line, "Neovim") {
			idx := strings.Index(line,"Neovim")

		diagnostics = append(diagnostics,lsp.Diagnostics{
			Range: LineRange(row, idx,idx+ len("Neovim")),
			Severity: 1,
			Source: "Please make sure we sure good language",
			Message:  "Please make sure we sure good language",
		} )
		}

	}

	return diagnostics
}

func (s *State) OpenDocument(document string, text string) []lsp.Diagnostics {
	s.Documents[document] = text

	return getDiagnosticsForFileText(text)
}
func (s *State) UpdateDocument(uri string, text string) []lsp.Diagnostics {
	s.Documents[uri] = text


	return getDiagnosticsForFileText(text)
}

func (s *State) Hover(id int, uri string, position lsp.Position) lsp.HoverResponse {
	 document := s.Documents[uri]
		return lsp.HoverResponse{
			Response: lsp.Response{
				RPC: "2.0",
				Id: &id,
			},
			Result: lsp.HoverResult{
				Contents: fmt.Sprintf("File: %s, Characters %d", uri, len(document)),
			},
		}
}
func (s *State) Definition(id int, uri string, position lsp.Position) lsp.DefinitionResponse {
		return lsp.DefinitionResponse{
			Response: lsp.Response{
				RPC: "2.0",
				Id: &id,
			},
			Result: lsp.Location{
				Uri: uri,
				Range: lsp.Range{
					Start: lsp.Position{
					Line: position.Line- 1,
					Character: 0,
					},	
					End: lsp.Position{
					Line: position.Line- 1,
					Character: 0,
					},	
				},
			},
		}
}

func (s *State) TextDocumentCompletion(id int, uri string, position lsp.Position, logger *log.Logger) lsp.CompletionResponse {

	items := []lsp.CompletionItem{}

	for i := 1; i <= 6; i++ {
		item := lsp.CompletionItem{
			Label: strings.Repeat("#", i),
			Detail: fmt.Sprintf("h%d", i),
			Documentation: fmt.Sprintf("This is a header of level %d", i),
			Kind: lsp.Keyword,
		}
		items = append(items, item)
	}
		items = append(items,  lsp.CompletionItem{
			Label: "*",
			Detail: "italic",
			Documentation: "this is italic",
			Kind: lsp.Keyword,
		}, lsp.CompletionItem{
			Label: "**",
			Detail: "bold",
			Documentation: "this is bold",
			Kind: lsp.Keyword,
		}, lsp.CompletionItem{
			Label: ">",
			Detail: "Blockquote",
			Documentation: "this is bold",
			Kind: lsp.Keyword,
		}, lsp.CompletionItem{
			Label: "-",
			Detail: "list Item",
			Documentation: "This is a list item",
			Kind: lsp.Keyword,
		}, lsp.CompletionItem{
			Label: "-",
			Detail: "list Item",
			Documentation: "This is a list item",
			Kind: lsp.Keyword,
		})

	return lsp.CompletionResponse{
		Response: lsp.Response{
			RPC: "2.0",
			Id: &id,
		},
		Result: items,
	}
}

func LineRange( line,  start, end int) lsp.Range{
	return lsp.Range{
		Start: lsp.Position{
			Line: line,
			Character: start,
		},
		End: lsp.Position{
			Line:  line,
			Character: end,
		},
	}

}