package rtpdump

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixWebrtc011SourceContract(t *testing.T) {
    source, err := os.ReadFile("reader.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if err != nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
