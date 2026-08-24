package webrtc

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixWebrtc015SourceContract(t *testing.T) {
    source, err := os.ReadFile("peerconnection.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if pc.configuration.ICECandidatePoolSize != configuration.ICECandidatePoolSize &&") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if pc.configuration.ICECandidatePoolSize != configuration.ICECandidatePoolSize ||") {
        t.Fatalf("mutated source contract is still present")
    }
}
