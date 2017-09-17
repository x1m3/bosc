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

func TestNewTree(t *testing.T) {
	for _, treeType := range bosc.TreeTypes() {
		_, err := bosc.NewTree(treeType)
		if err!=nil {
			t.Errorf("Error creating tree <%v>", err)
		}
	}
}



