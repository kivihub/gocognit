package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestCognitive(t *testing.T) {
	os.Args = []string{"gocognit", "-over=-1", "../../testdata"}
	CaptureStdout(func(out string) {
		if out == "" {
			t.Error("Expect not empty")
		}
	})
}

func TestCognitiveIgnore(t *testing.T) {
	os.Args = []string{"gocognit", "-over=-1", "-ignoreContent=(// Code generated by .*)|(// Autogenerated by .*)", "../../testdata"}
	CaptureStdout(func(out string) {
		if strings.Contains(out, "d1.go") || strings.Contains(out, "d2.go") {
			t.Error("Expect ignore d.go")
		}
	})
}

func CaptureStdout(processOutput func(string)) {
	endCapture := doCaptureStdout()
	main()
	capturedContent := endCapture()
	fmt.Print("Captured: ", capturedContent)
	processOutput(capturedContent)
}

// doCaptureStdout 目前只适合拦截少量输出，如果过大超过缓冲区，则会阻塞
func doCaptureStdout() func() string {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	return func() string {
		w.Close()
		var buf bytes.Buffer
		buf.ReadFrom(r)
		os.Stdout = oldStdout
		return buf.String()
	}
}
