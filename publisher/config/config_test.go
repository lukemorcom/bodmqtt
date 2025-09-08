package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	tmpDir := t.TempDir()
	yamlPath := filepath.Join(tmpDir, "config-testing.yml")

	yamlData := `apis:
  - name: local
    url: "http://example.local"
    strategy: ping
events:
  - name: ping_local
    api: local
    topic: test/topic`

	// write the dummy yaml file for LoadConfig to read
	if err := os.WriteFile(yamlPath, []byte(yamlData), 0644); err != nil {
		t.Fatal(err)
	}

	_, err := LoadConfig(yamlPath)
	if err != nil {
		t.Fatal(err)
	}
}
