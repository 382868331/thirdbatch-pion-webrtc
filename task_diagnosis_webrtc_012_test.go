package webrtc

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisWebrtc012SourceContract(t *testing.T) {
    source, err := os.ReadFile("mediaengine.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "return true") {
        t.Fatalf("expected source contract is missing")
    }
}
