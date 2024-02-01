// config_test.go
package main

import (
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		name      string
		confPath  string
		setEnv    bool
		writeFile bool
		wantError bool
	}{
		{
			name:      "Env Variable Exists",
			confPath:  "",
			setEnv:    true,
			writeFile: false,
			wantError: false,
		},
		{
			name:      "Config File Exists",
			setEnv:    false,
			writeFile: true,
			wantError: false,
		},
		{
			name:      "No Config Found",
			confPath:  "",
			setEnv:    false,
			writeFile: false,
			wantError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.setEnv {
				os.Setenv("OPENAI_API_KEY", "test-key")
				defer os.Unsetenv("OPENAI_API_KEY")
			}

			if tc.writeFile {
				tempFile, err := os.CreateTemp("", "test_config.*.conf")
				if err != nil {
					t.Fatal(err)
				}
				defer os.Remove(tempFile.Name())

				_, err = tempFile.WriteString("OPENAI_API_KEY=test-key")
				if err != nil {
					t.Fatal(err)
				}
				tempFile.Close()

				tc.confPath = tempFile.Name()
			}

			_, err := NewConfig(tc.confPath)

			if (err != nil) != tc.wantError {
				t.Errorf("NewConfig() error = %v, wantError %v", err, tc.wantError)
			}
		})
	}
}
