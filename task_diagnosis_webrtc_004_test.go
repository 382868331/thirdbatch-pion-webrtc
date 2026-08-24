package webrtc

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisWebrtc004SourceContract(t *testing.T) {
    source, err := os.ReadFile("peerconnection.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if handler, ok := pc.onICEConnectionStateChangeHandler.Load().(func(ICEConnectionState)); ok && handler != nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
