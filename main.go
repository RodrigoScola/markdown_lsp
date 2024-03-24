package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"mylsp/pkg/analysis"
	"mylsp/pkg/lsp"
	"mylsp/pkg/rpc"
	"os"
)

func main() {
	logger:= getLogger("D:/code/lsp/log.txt")
	logger.Println("hey i started")
	state := analysis.NewState()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	writer := os.Stdout

	for scanner.Scan() {
		msg := scanner.Bytes()

		method, content, err :=rpc.DecodeMessage(msg)

		if err != nil {
			 logger.Printf("ayo, we got this eror %s", err)
		  return
		}

		handleMessage(method, content,&state,writer, logger)
	}
}

func handleMessage(method string , content []byte ,state *analysis.State,writer io.Writer, logger *log.Logger) {
	logger.Printf("Received message with method: %s", method)
	switch method {
	case "initialize": 

		var request lsp.InitializeRequest;
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("we couldnt parse this %s", err)
		}

		logger.Printf("connected to: %s %s", 
		request.Params.ClientInfo.Name,
		request.Params.ClientInfo.Version)

		msg := lsp.NewInitializeResponse(request.ID)



		writeResponse(writer, msg)
		logger.Println("send the response")

	case "textDocument/didOpen": 
		var req lsp.DidOpenTextDocumentNotification;
		if err := json.Unmarshal(content, &req); err != nil {
			logger.Printf("textDocument/didOpen %s", err)
			return 
		}

		state.OpenDocument(req.Params.TextDocument.Uri, req.Params.TextDocument.Text)

		logger.Printf("Opened: %s ", req.Params.TextDocument.Uri) 

	case "textDocument/didChange": 
		var req lsp.TextDocumentDidChangeNotification;
		if err := json.Unmarshal(content, &req); err != nil {
			logger.Printf("textDocument/didChange %s", err)
			return 
		}


		logger.Printf("Changed %s ", req.Params.TextDocument.Uri)

		for _, change := range req.Params.ContentChanges {
		diagnostics :=		state.UpdateDocument(req.Params.TextDocument.Uri, change.Text)

		writeResponse(writer, lsp.PublishDiagnosticsNotification{
			Notification: lsp.Notification{
				Method: "textDocument/publishDiagnostics",
				RPC: "2.0",
			},
			Params: lsp.PublishDiagnosticsParams{
				Uri: req.Params.TextDocument.Uri,
				Diagnostics: diagnostics,
			},
		})
		}

	case "textDocument/hover": 
		var req lsp.HoverRequest;
		if err := json.Unmarshal(content, &req); err != nil {
			logger.Printf("textDocument/hover %s", err)
			return 
		}
		respose := state.Hover(req.ID, req.Params.TextDocument.Uri, req.Params.Position)

		writeResponse(writer, respose)
	case "textDocument/definition": 
		var req lsp.DefinitionRequest;
		if err := json.Unmarshal(content, &req); err != nil {
			logger.Printf("textDocument/definition %s", err)
			return 
		}

		respose := state.Definition(req.ID, req.Params.TextDocument.Uri, req.Params.Position)
		writeResponse(writer, respose)

	case "textDocument/completion": 
		var req lsp.CompletionRequest;
		if err := json.Unmarshal(content, &req); err != nil {
			logger.Printf("textDocument/definition %s", err)
			return 
		}

		respose := state.TextDocumentCompletion(req.ID, req.Params.TextDocument.Uri, req.Params.Position, logger)
		writeResponse(writer, respose)

	}
}

func writeResponse(writer io.Writer,msg any) {
		reply := rpc.EncodeMessage(msg)
		writer.Write([]byte(reply))
}

func getLogger(filename string) *log.Logger {
	logfile ,err:= os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)		

	if err != nil {
		panic("hey you didnt give me a good file")
	}

	return log.New(logfile, "[lsp]", log.Ldate|log.LUTC|log.Lshortfile)
}