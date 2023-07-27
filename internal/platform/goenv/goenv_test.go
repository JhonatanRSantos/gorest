package goenv

import (
	"testing"
)

func TestLoadString(t *testing.T) {
	t.Setenv("TEST_APP_STRING", "STRING_VALUE")
	if Load("TEST_APP_STRING", "DEFAULT_STRING_VALUE") != "STRING_VALUE" {
		t.Fatal("failed to load env var (string)")
	}
}

func TestLoadInt(t *testing.T) {
	t.Setenv("TEST_APP_INT", "1")
	value := Load("TEST_APP_INT", int64(-1))
	if value != 1 {
		t.Fatal("failed to load env var (int64)")
	}
}

func TestLoadFloat(t *testing.T) {
	t.Setenv("TEST_APP_FLOAT", "1.75")
	value := Load("TEST_APP_FLOAT", float64(-1))
	if value != 1.75 {
		t.Fatal("failed to load env var (float64)")
	}
}

func TestLoadBool(t *testing.T) {
	t.Setenv("TEST_APP_BOOL", "true")
	value := Load("TEST_APP_BOOL", false)
	if value != true {
		t.Fatal("failed to load env var (bool)")
	}
}
