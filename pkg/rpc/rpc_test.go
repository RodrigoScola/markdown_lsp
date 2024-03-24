package rpc_test

import (
	"mylsp/pkg/rpc"
	"testing"
)

type EncodingExample struct {
	Testing bool
}
func TestEncode(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodeMessage(EncodingExample{Testing: true})

	if expected != actual {
		t.Fatalf("Expected: %s, actual: %s", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	expected := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	method, content, err := rpc.DecodeMessage([]byte(expected))
	contentLenght := len(content)

	if (err != nil) {
		t.Fatal(err)
	}

	if contentLenght != 15 {
		t.Fatalf("Expected 15, got : %d", contentLenght)
	}
	if method != "hi" {
		t.Fatalf("expected: \"hi\" got %s", method)
	}
}