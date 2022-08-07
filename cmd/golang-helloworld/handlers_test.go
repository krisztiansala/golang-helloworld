package main

import (
	"bytes"
	"testing"
)

func TestHelloMessage(t *testing.T) {
	testCases := map[string]struct {
		name string
		want string
	}{
		"empty": {
			name: "",
			want: "Hello Stranger",
		},
		"basic": {
			name: "Krisztian",
			want: "Hello Krisztian",
		},
		"camelCase": {
			name: "KrisztianSala",
			want: "Hello Krisztian Sala",
		},
		"numbers": {
			name: "1Kriszt2ian3Sala4",
			want: "Hello Krisztian Sala",
		},
		"unicode": {
			name: "Krisztián ",
			want: "Hello Krisztián",
		},
		"specialChars": {
			name: "123%^$:?>''",
			want: "Hello Stranger",
		},
	}
	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			if got := helloMessage(testCase.name); got != testCase.want {
				t.Errorf("helloMessage(%v) got %v, want %v", testCase.name, got, testCase.want)
			}
		})
	}
}
func TestVersionResponse(t *testing.T) {
	testCases := map[string]struct {
		hash    string
		project string
		want    []byte
	}{
		"empty": {
			hash:    "",
			project: "",
			want:    []byte(`{"hash":"","project":""}`),
		},
		"hashOnly": {
			hash:    "123",
			project: "",
			want:    []byte(`{"hash":"123","project":""}`),
		},
		"projectOnly": {
			hash:    "",
			project: "my-project",
			want:    []byte(`{"hash":"","project":"my-project"}`),
		},
		"complete": {
			hash:    "asd123",
			project: "my-project",
			want:    []byte(`{"hash":"asd123","project":"my-project"}`),
		},
	}
	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			if got := versionResponse(testCase.hash, testCase.project); !bytes.Equal(got, testCase.want) {
				t.Errorf("Version response %v does not match wanted value %v", string(got), string(testCase.want))
			}
		})
	}
}
