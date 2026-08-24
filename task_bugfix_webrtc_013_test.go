package webrtc

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixWebrtc013SourceContract(t *testing.T) {
    source, err := os.ReadFile("sdp.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if split[0] == sdp.SemanticTokenFlowIdentification { //nolint:nestif") {
        t.Fatalf("expected source contract is missing")
    }
}
