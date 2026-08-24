package webrtc

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisWebrtc020SourceContract(t *testing.T) {
    source, err := os.ReadFile("certificate.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if cert == nil || privateKey == nil {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if cert == nil && privateKey == nil {") {
        t.Fatalf("mutated source contract is still present")
    }
}
