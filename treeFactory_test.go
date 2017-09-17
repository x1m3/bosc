package bosc_test

import (
	"testing"
	"bosc"
)

func TestNewTreeDoesntExists(t *testing.T) {
	_, err := bosc.NewTree("doesnt_exists")
	if err == nil {
		t.Error("Error getting a tree type that doesn't exists")
	}
}


