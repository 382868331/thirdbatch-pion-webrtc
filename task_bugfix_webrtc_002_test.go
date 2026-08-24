package webrtc

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixWebrtc002SourceContract(t *testing.T) {
    source, err := os.ReadFile("rtptransceiver.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if prevSender := t.Sender(); prevSender != nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if prevSender := t.Sender(); prevSender == nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
