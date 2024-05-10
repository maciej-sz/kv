package pkg

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParseKeyValueFile_Success(t *testing.T) {
	// Create a temporary file with test data
	content := `key1=value1
key2="value2"
# This is a comment
key3='value3'
`
	tmpFile, err := os.CreateTemp(os.TempDir(), "testfile-*.txt")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())
	_, err = tmpFile.WriteString(content)
	assert.NoError(t, err)
	tmpFile.Close()

	// Call the function
	result, err := ParseKeyValueFile(tmpFile.Name())
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "value1", result["key1"].Val)
	assert.Equal(t, "\"", result["key2"].Quote)
	assert.Equal(t, "value2", result["key2"].Val)
	assert.Equal(t, "'", result["key3"].Quote)
	assert.Equal(t, "value3", result["key3"].Val)
}

func TestParseKeyValueFile_Failure(t *testing.T) {
	// Create a temporary file with test data
	content := `key1=value1
key2
key3='value3'
`
	tmpFile, err := os.CreateTemp(os.TempDir(), "testfile-*.txt")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())
	_, err = tmpFile.WriteString(content)
	assert.NoError(t, err)
	tmpFile.Close()

	// Call the function
	_, err = ParseKeyValueFile(tmpFile.Name())
	assert.Error(t, err)
	assert.Equal(t, "invalid line: key2", err.Error())
}
