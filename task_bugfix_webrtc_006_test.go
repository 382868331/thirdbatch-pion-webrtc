package webrtc

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixWebrtc006SourceContract(t *testing.T) {
    source, err := os.ReadFile("sdp.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "mediaDescr := parsed.MediaDescriptions[0]") {
        t.Fatalf("expected source contract is missing")
    }
}
