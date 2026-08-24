package webrtc

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixWebrtc018SourceContract(t *testing.T) {
    source, err := os.ReadFile("icegatherer.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "rules := make([]ice.AddressRewriteRule, 0, len(ips)+1)") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "rules := make([]ice.AddressRewriteRule, 0, len(ips)- 1)") {
        t.Fatalf("mutated source contract is still present")
    }
}
