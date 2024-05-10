package pkg

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestSaveToFile_Success(t *testing.T) {
	tmpFile, err := os.CreateTemp(os.TempDir(), "testfile-*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up after the test

	// Test data
	data := map[string]*Value{
		"key1": {"value1", "\""},
		"key2": {"value2", "'"},
		"key3": {"value3", ""},
	}

	assert.NoError(t, SaveToFile(tmpFile.Name(), data))

	content, err := os.ReadFile(tmpFile.Name())
	assert.NoError(t, err)

	contentStr := string(content)
	assert.True(t, strings.Contains(contentStr, "key1=\"value1\"\n"))
	assert.True(t, strings.Contains(contentStr, "key2='value2'\n"))
	assert.True(t, strings.Contains(contentStr, "key3=value3\n"))
}

func TestSaveToFile_FailOnFileCreate(t *testing.T) {
	invalidFilePath := "/invalid/path/to/file.txt"
	data := map[string]*Value{"key1": {"\"", "value1"}}

	err := SaveToFile(invalidFilePath, data)
	assert.Error(t, err)
}

func TestSaveToFile_EmptyData(t *testing.T) {
	tmpFile, err := os.CreateTemp(os.TempDir(), "testfile-*.txt")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	data := make(map[string]*Value) // Empty map

	assert.NoError(t, SaveToFile(tmpFile.Name(), data))

	content, err := os.ReadFile(tmpFile.Name())
	assert.NoError(t, err)

	assert.True(t, len(content) == 0)
}
