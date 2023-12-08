package tests

import (
	"os"
	"testing"
)

// Global setup and teardown
func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}
