package cmd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

func TestGetValue_Success(t *testing.T) {

	var tmpFile *os.File
	var err error

	//<editor-fold desc="Setup test">

	// Create a temporary file with test data
	content := `key1=value1
key2="value2"
# This is a comment
key3='value3'
`
	tmpFile, err = os.CreateTemp(os.TempDir(), "testfile-*.txt")
	assert.NoError(t, err)

	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Errorf("failed to remove file %s: %v", name, err)
		}
	}(tmpFile.Name())

	_, err = tmpFile.WriteString(content)
	assert.NoError(t, err)

	err = tmpFile.Close()
	assert.NoError(t, err)
	//</editor-fold>

	testCases := []struct {
		name          string
		key           string
		defaultValue  string
		expectedValue string
		expectedError string
	}{
		{
			name:          "Simple key",
			key:           "key1",
			expectedValue: "value1",
			expectedError: "",
		},
		{
			name:          "Key with double quotes",
			key:           "key2",
			expectedValue: "value2",
			expectedError: "",
		},
		{
			name:          "Key with single quotes",
			key:           "key3",
			expectedValue: "value3",
			expectedError: "",
		},
		{
			name:          "Not found key",
			key:           "bogus-key",
			expectedValue: "",
			expectedError: "key doesn't exist",
		},
		{
			name:          "Not found key with default",
			key:           "bogus-key",
			defaultValue:  "default-value",
			expectedValue: "default-value",
			expectedError: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cmd := newGetCmd()
			outBuff := bytes.NewBufferString("")
			errBuff := bytes.NewBufferString("")
			cmd.SetOut(outBuff)
			cmd.SetErr(errBuff)
			cmd.SetArgs([]string{tmpFile.Name(), tc.key})

			err = cmd.Execute()
			assert.NoError(t, err)

			outContent, err := io.ReadAll(outBuff)
			assert.NoError(t, err)
			errContent, err := io.ReadAll(errBuff)
			assert.NoError(t, err)

			assert.Equal(t, tc.expectedValue, string(outContent))
			assert.Equal(t, tc.expectedError, string(errContent))
		})
	}
}
