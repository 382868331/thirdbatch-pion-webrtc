package webrtc

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixWebrtc003SourceContract(t *testing.T) {
    source, err := os.ReadFile("peerconnection.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "(dtlsTransportState == DTLSTransportStateNew || dtlsTransportState == DTLSTransportStateClosed):") {
        t.Fatalf("expected source contract is missing")
    }
}
