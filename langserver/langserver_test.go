// Copyright 2019 Tobias Guggenmos
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package langserver

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/prometheus-community/promql-langserver/config"
	"github.com/prometheus-community/promql-langserver/internal/vendored/go-tools/jsonrpc2"
	"github.com/prometheus-community/promql-langserver/internal/vendored/go-tools/lsp/protocol"
)

// TestNotImplemented checks whether unimplemented functions return the approbiate Error.
func TestNotImplemented(*testing.T) { //nolint: gocognit, funlen, gocyclo
	s := &server{}

	err := s.DidChangeWorkspaceFolders(context.Background(), &protocol.DidChangeWorkspaceFoldersParams{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	err = s.DidSave(context.Background(), &protocol.DidSaveTextDocumentParams{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	err = s.WillSave(context.Background(), &protocol.WillSaveTextDocumentParams{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	err = s.DidChangeWatchedFiles(context.Background(), &protocol.DidChangeWatchedFilesParams{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	err = s.Progress(context.Background(), &protocol.ProgressParams{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.SelectionRange(context.Background(), &protocol.SelectionRangeParams{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	err = s.SetTraceNotification(context.Background(), &protocol.SetTraceParams{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	err = s.LogTraceNotification(context.Background(), &protocol.LogTraceParams{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.Implementation(context.Background(), &protocol.ImplementationParams{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.TypeDefinition(context.Background(), &protocol.TypeDefinitionParams{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.DocumentColor(context.Background(), &protocol.DocumentColorParams{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.ColorPresentation(context.Background(), &protocol.ColorPresentationParams{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.FoldingRange(context.Background(), &protocol.FoldingRangeParams{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.NonstandardRequest(context.Background(), "", nil)
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.Declaration(context.Background(), &protocol.DeclarationParams{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.WillSaveWaitUntil(context.Background(), &protocol.WillSaveTextDocumentParams{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.Resolve(context.Background(), &protocol.CompletionItem{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.Definition(context.Background(), &protocol.DefinitionParams{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.References(context.Background(), &protocol.ReferenceParams{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.DocumentHighlight(context.Background(), &protocol.DocumentHighlightParams{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.DocumentSymbol(context.Background(), &protocol.DocumentSymbolParams{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.CodeAction(context.Background(), &protocol.CodeActionParams{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.Symbol(context.Background(), &protocol.WorkspaceSymbolParams{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.CodeLens(context.Background(), &protocol.CodeLensParams{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.ResolveCodeLens(context.Background(), &protocol.CodeLens{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.Formatting(context.Background(), &protocol.DocumentFormattingParams{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.RangeFormatting(context.Background(), &protocol.DocumentRangeFormattingParams{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.OnTypeFormatting(context.Background(), &protocol.DocumentOnTypeFormattingParams{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.Rename(context.Background(), &protocol.RenameParams{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.PrepareRename(context.Background(), &protocol.PrepareRenameParams{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.DocumentLink(context.Background(), &protocol.DocumentLinkParams{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.ResolveDocumentLink(context.Background(), &protocol.DocumentLink{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.ExecuteCommand(context.Background(), &protocol.ExecuteCommandParams{})
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.IncomingCalls(context.Background(), nil)
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.OutgoingCalls(context.Background(), nil)
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.PrepareCallHierarchy(context.Background(), nil)
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.SemanticTokens(context.Background(), nil)
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.SemanticTokensEdits(context.Background(), nil)
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	_, err = s.SemanticTokensRange(context.Background(), nil)
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	err = s.WorkDoneProgressCancel(context.Background(), nil)
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}

	err = s.WorkDoneProgressCreate(context.Background(), nil)
	if err != nil && err.(*jsonrpc2.Error).Code != jsonrpc2.CodeMethodNotFound {
		panic("Expected a jsonrpc2 Error with CodeMethodNotFound")
	}
}

// dummyStream is a fake jsonrpc2.Stream for Test purposes.
type dummyStream struct {
	readQueue []byte
}

func (d *dummyStream) Read(_ context.Context) ([]byte, int64, error) {
	ret := d.readQueue
	d.readQueue = []byte{}

	return ret, int64(len(ret)), nil
}

func (d *dummyStream) Write(_ context.Context, text []byte) (int64, error) {
	return int64(len(text)), nil
}

// Push adds a text to the readQueue.
func (d *dummyStream) Push(text []byte) {
	d.readQueue = append(d.readQueue, text...)
}

type dummyWriter struct{}

func (d *dummyWriter) Write(text []byte) (int, error) {
	return len(text), nil
}

// TestServerState tries to emulate a full server lifetime.
func TestServer(t *testing.T) { //nolint:funlen, gocognit, gocyclo
	var stream jsonrpc2.Stream = &dummyStream{}
	stream = jSONLogStream(stream, &dummyWriter{})
	_, server := ServerFromStream(context.Background(), stream, &config.Config{LogFormat: config.TextFormat})
	s := server.server

	// Initialize Server
	_, err := s.Initialize(context.Background(), &protocol.ParamInitialize{})
	if err != nil {
		panic("Failed to initialize Server")
	}

	_, err = s.Initialize(context.Background(), &protocol.ParamInitialize{})
	if err == nil {
		panic("cannot initialize server twice")
	}
	// Confirm Initialisation
	err = s.Initialized(context.Background(), &protocol.InitializedParams{})
	if err != nil {
		panic("Failed to initialize Server")
	}

	err = s.Initialized(context.Background(), &protocol.InitializedParams{})
	if err == nil {
		panic("cannot confirm server initialisation twice")
	}

	// Add a document to the server
	err = s.DidOpen(context.Background(), &protocol.DidOpenTextDocumentParams{
		TextDocument: protocol.TextDocumentItem{
			URI:        "test.promql",
			LanguageID: "promql",
			Version:    0,
			Text:       "",
		},
	})
	if err != nil {
		panic("Failed to open document")
	}

	// Apply a Full Change to the document
	err = s.DidChange(context.Background(), &protocol.DidChangeTextDocumentParams{
		TextDocument: protocol.VersionedTextDocumentIdentifier{
			Version: 2,
			TextDocumentIdentifier: protocol.TextDocumentIdentifier{
				URI: "test.promql",
			},
		},
		ContentChanges: []protocol.TextDocumentContentChangeEvent{
			{
				Range:       nil,
				RangeLength: 0,
				Text:        "sum()",
			},
		},
	})
	if err != nil {
		panic("Failed to apply full change to document")
	}

	hover, err := s.Hover(context.Background(), &protocol.HoverParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{
				URI: "test.promql",
			},
			Position: protocol.Position{
				Line:      0.0,
				Character: 1.0,
			},
		},
	})

	if err != nil {
		panic("Failed to get hovertext")
	}

	if hover == nil || !strings.Contains(hover.Contents.Value, "sum") {
		fmt.Println(hover)
		panic("unexpected or no hovertext")
	}

	// Apply a Full Change to the document
	err = s.DidChange(context.Background(), &protocol.DidChangeTextDocumentParams{
		TextDocument: protocol.VersionedTextDocumentIdentifier{
			Version: 3,
			TextDocumentIdentifier: protocol.TextDocumentIdentifier{
				URI: "test.promql",
			},
		},
		ContentChanges: []protocol.TextDocumentContentChangeEvent{
			{
				Range:       nil,
				RangeLength: 0,
				Text:        "metric_name",
			},
		},
	})
	if err != nil {
		panic("Failed to apply full change to document")
	}

	hover, err = s.Hover(context.Background(), &protocol.HoverParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{
				URI: "test.promql",
			},
			Position: protocol.Position{
				Line:      0.0,
				Character: 1.0,
			},
		},
	})

	if err != nil {
		panic("Failed to get hovertext")
	}

	if hover == nil || !strings.Contains(hover.Contents.Value, "metric_name") {
		fmt.Println(hover)
		panic("unexpected or no hovertext")
	}
	// Apply a partial Change to the document
	err = s.DidChange(context.Background(), &protocol.DidChangeTextDocumentParams{
		TextDocument: protocol.VersionedTextDocumentIdentifier{
			Version: 4,
			TextDocumentIdentifier: protocol.TextDocumentIdentifier{
				URI: "test.promql",
			},
		},
		ContentChanges: []protocol.TextDocumentContentChangeEvent{
			{
				Range: &protocol.Range{
					Start: protocol.Position{
						Line:      0.0,
						Character: 0.0,
					},
					End: protocol.Position{
						Line:      0.0,
						Character: 0.0,
					},
				},
				RangeLength: 5,
				Text:        "rate(",
			},
		},
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to apply change to document: %s", err.Error()))
	}

	// Wait for diagnostics
	doc, err := s.cache.GetDocument("test.promql")
	if err != nil {
		panic("Failed to get document")
	}

	if diagnostics, err := doc.GetDiagnostics(); err != nil && len(diagnostics) != 0 {
		panic("expected nonempty diagnostics")
	}

	// Apply a partial Change to the document
	err = s.DidChange(context.Background(), &protocol.DidChangeTextDocumentParams{
		TextDocument: protocol.VersionedTextDocumentIdentifier{
			Version: 5,
			TextDocumentIdentifier: protocol.TextDocumentIdentifier{
				URI: "test.promql",
			},
		},
		ContentChanges: []protocol.TextDocumentContentChangeEvent{
			{
				Range: &protocol.Range{
					Start: protocol.Position{
						Line:      0.0,
						Character: 11.0,
					},
					End: protocol.Position{
						Line:      0.0,
						Character: 16.0,
					},
				},
				RangeLength: 5,
				Text:        ")",
			},
		},
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to apply change to document: %s", err.Error()))
	}

	// Wait for diagnostics
	doc, err = s.cache.GetDocument("test.promql")
	if err != nil {
		panic("Failed to get document")
	}

	if diagnostics, err := doc.GetDiagnostics(); err != nil && len(diagnostics) != 0 {
		panic("expected empty diagnostics")
	}

	var content string

	content, err = doc.GetContent()
	if err != nil {
		panic("failed to get document content")
	}

	if content != "rate(metric)" {
		panic(fmt.Sprintf("unexpected content, expected \"rate(metric)\", got %s", content))
	}

	// Apply a Full Change to the document
	err = s.DidChange(context.Background(), &protocol.DidChangeTextDocumentParams{
		TextDocument: protocol.VersionedTextDocumentIdentifier{
			Version: 6,
			TextDocumentIdentifier: protocol.TextDocumentIdentifier{
				URI: "test.promql",
			},
		},
		ContentChanges: []protocol.TextDocumentContentChangeEvent{
			{
				Range:       nil,
				RangeLength: 0,
				Text:        "rat",
			},
		},
	})
	if err != nil {
		panic("Failed to apply full change to document")
	}

	completion, err := s.Completion(context.Background(), &protocol.CompletionParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{
				URI: "test.promql",
			},
			Position: protocol.Position{
				Line:      0.0,
				Character: 1.0,
			},
		},
	})

	if err != nil || completion == nil || len(completion.Items) == 0 || completion.Items[0].Label != "rate" {
		fmt.Println(completion)
		panic("Failed to get completion")
	}

	// Apply a Full Change to the document
	err = s.DidChange(context.Background(), &protocol.DidChangeTextDocumentParams{
		TextDocument: protocol.VersionedTextDocumentIdentifier{
			Version: 7,
			TextDocumentIdentifier: protocol.TextDocumentIdentifier{
				URI: "test.promql",
			},
		},
		ContentChanges: []protocol.TextDocumentContentChangeEvent{
			{
				Range:       nil,
				RangeLength: 0,
				Text:        "rat()",
			},
		},
	})
	if err != nil {
		panic("Failed to apply full change to document")
	}

	completion, err = s.Completion(context.Background(), &protocol.CompletionParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{
				URI: "test.promql",
			},
			Position: protocol.Position{
				Line:      0.0,
				Character: 1.0,
			},
		},
	})

	if err != nil || completion == nil || len(completion.Items) == 0 || completion.Items[1].Label != "rate" {
		fmt.Println(completion)
		panic("Failed to get completion")
	}

	// Close a document
	err = s.DidClose(context.Background(), &protocol.DidCloseTextDocumentParams{
		TextDocument: protocol.TextDocumentIdentifier{
			URI: "test.promql",
		},
	})
	if err != nil {
		panic("Failed to close document")
	}

	_, err = s.cache.GetDocument("test.promql")
	if err == nil {
		panic("getting a closed document should have failed")
	}

	// Close a document twice
	err = s.DidClose(context.Background(), &protocol.DidCloseTextDocumentParams{
		TextDocument: protocol.TextDocumentIdentifier{
			URI: "test.promql",
		},
	})
	if err == nil {
		panic("should have failed to close document")
	}

	// Reopen a closed document
	err = s.DidOpen(context.Background(), &protocol.DidOpenTextDocumentParams{
		TextDocument: protocol.TextDocumentItem{
			URI:        "test.promql",
			LanguageID: "promql",
			Version:    0,
			Text:       "abs()",
		},
	})
	if err != nil {
		panic("Failed to reopen document")
	}

	signature, err := s.SignatureHelp(context.Background(), &protocol.SignatureHelpParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{
				URI: "test.promql",
			},
			Position: protocol.Position{
				Line:      1.0,
				Character: 0.0,
			},
		},
	})

	if err != nil {
		panic("Failed to get signature")
	}

	if signature != nil && len(signature.Signatures) != 0 {
		fmt.Println(signature)
		panic("Wrong number of signatures returned")
	}

	signature, err = s.SignatureHelp(context.Background(), &protocol.SignatureHelpParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{
				URI: "test.promql",
			},
			Position: protocol.Position{
				Line:      0,
				Character: 4,
			},
		},
	})

	if err != nil {
		panic("Failed to get signature")
	}

	if signature == nil || len(signature.Signatures) != 1 {
		fmt.Println(signature.Signatures)
		panic("Wrong number of signatures returned")
	}

	hover, err = s.Hover(context.Background(), &protocol.HoverParams{
		TextDocumentPositionParams: protocol.TextDocumentPositionParams{
			TextDocument: protocol.TextDocumentIdentifier{
				URI: "test.promql",
			},
			Position: protocol.Position{
				Line:      0.0,
				Character: 1.0,
			},
		},
	})

	if err != nil {
		panic("Failed to get hovertext")
	}

	if hover == nil || !strings.Contains(hover.Contents.Value, "abs") {
		fmt.Println(hover)
		panic("unexpected or no hovertext")
	}

	// Shutdown Server
	err = s.Shutdown(context.Background())
	if err != nil {
		panic("Failed to initialize Server")
	}

	err = s.Shutdown(context.Background())
	if err == nil {
		panic("cannot shutdown server twice")
	}
	// Left out until it does something else than calling os.Exit()
	// Confirm Shutdown
	err = s.Exit(context.Background())
	if err != nil {
		panic("Failed to initialize Server")
	}
}
