package log

import (
	"bytes"
	"github.com/MuratSs/assert"
	"net"
	"strings"
	"sync"
	"testing"
)

func TestSetLevel_WithWarning(t *testing.T) {
	// Since `go test` doesn't reset the test context for each unit test,
	// all subsequent tests should emit only warning and error messages.
	SetLevel(WARN)
}

func TestEmit_WithInfo_ShouldBeEmpty(t *testing.T) {
	Assert := assert.With(t)

	var buf bytes.Buffer
	SetWriter(&buf)

	Info("test")

	Assert.That(buf.String()).IsEmpty()
}

func TestEmit_WithWarn_ShouldNotBeEmpty(t *testing.T) {
	Assert := assert.With(t)

	var buf bytes.Buffer
	SetWriter(&buf)

	Warn("test")

	Assert.That(buf.String()).IsNotEmpty()
}

func TestEmit_WithWarnAndFields_ShouldNotBeEmpty(t *testing.T) {
	Assert := assert.With(t)

	var buf bytes.Buffer
	SetWriter(&buf)

	Warn("test", Int16("port", 9000), String("hostname", "localhost"))

	s := buf.String()
	Assert.That(s).IsNotEmpty()

	if !strings.Contains(s, "port=9000") {
		t.Fatal("Port field was not in the log message")
	}

	if !strings.Contains(s, "hostname=localhost") {
		t.Fatal("Hostname field was not in the log message")
	}
}

func TestSetServer_WithServer_ShouldReceive(t *testing.T) {
	Assert := assert.With(t)

	// Create a simple UDP server that exits after one message is received.
	addr, err := net.ResolveUDPAddr("udp", "localhost:9999")
	Assert.That(err).IsOk()

	sock, err := net.ListenUDP("udp", addr)
	Assert.That(err).IsOk()

	defer sock.Close()

	// Waits for the server goroutine to finish.
	var wait sync.WaitGroup

	go func() {
		buf := make([]byte, 4096)
		len, _, err := sock.ReadFromUDP(buf)
		Assert.That(err).IsOk()
		Assert.That(len).IsGreaterThan(0)
		Assert.That(buf).IsNotNil()

		s := string(buf)
		Assert.That(s).IsNotEmpty()

		if !strings.Contains(s, "\"port\":9000") {
			t.Fatal("Port field was not in the log message")
		}

		if !strings.Contains(s, "\"hostname\":\"localhost\"") {
			t.Fatal("Hostname field was not in the log message")
		}

		// Signal that the test is complete.
		wait.Done()
	}()

	// Prevent output on the console.
	var buf bytes.Buffer
	SetWriter(&buf)

	// Send a message to the server.
	SetServer("localhost:9999")
	wait.Add(1)
	Warn("test", Int16("port", 9000), String("hostname", "localhost"))

	// Wait for the test to complete.
	wait.Wait()
}
