package utils

import (
	"os"
	"testing"
)

func TestLoadEnv(t *testing.T) {
	testCases := []struct {
		fileContent string
		envVars     map[string]string
		expectedErr bool
	}{
		{
			fileContent: "VAR1=value1\nVAR2=value2\nVAR3=value3\n",
			envVars: map[string]string{
				"VAR1": "value1",
				"VAR2": "value2",
				"VAR3": "value3",
			},
			expectedErr: false,
		},
		{
			fileContent: "VAR1=value1\n#VAR2=value2\nVAR3=value3\n",
			envVars: map[string]string{
				"VAR1": "value1",
				"VAR3": "value3",
			},
			expectedErr: false,
		},
		{
			fileContent: "VAR1=value1\nVAR2=value2\nVAR3=value3",
			envVars: map[string]string{
				"VAR1": "value1",
				"VAR2": "value2",
				"VAR3": "value3",
			},
			expectedErr: false,
		},
	}

	for _, tc := range testCases {
		tmpfile, err := os.CreateTemp("", "")
		if err != nil {
			t.Fatal(err)
		}

		_, err = tmpfile.WriteString(tc.fileContent)
		if err != nil {
			t.Fatal(err)
		}

		if err := loadEnv(tmpfile.Name()); (err != nil) != tc.expectedErr {
			t.Errorf("loadEnv(%q) = %v; expected error? %v", tmpfile.Name(), err, tc.expectedErr)
		}

		for key, expectedValue := range tc.envVars {
			if actualValue := os.Getenv(key); actualValue != expectedValue {
				t.Errorf("os.Getenv(%q) = %q; expected %q", key, actualValue, expectedValue)
			}
		}
		tmpfile.Close()
	}
}
