package webrtc

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisWebrtc016SourceContract(t *testing.T) {
    source, err := os.ReadFile("icetransport.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if role == nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
