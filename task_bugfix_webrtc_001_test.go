package webrtc

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixWebrtc001SourceContract(t *testing.T) {
    source, err := os.ReadFile("mediaengine.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if _, taken := m.negotiatedHeaderExtensions[id]; idAvailable && !taken {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if _, taken := m.negotiatedHeaderExtensions[id]; idAvailable || !taken {") {
        t.Fatalf("mutated source contract is still present")
    }
}
