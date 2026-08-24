package webrtc

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixWebrtc007SourceContract(t *testing.T) {
    source, err := os.ReadFile("rtpsender.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if encoding.track.RID() == track.RID() {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if encoding.track.RID() != track.RID() {") {
        t.Fatalf("mutated source contract is still present")
    }
}
