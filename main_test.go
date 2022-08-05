package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseFlagsCorrect(t *testing.T) {
	var tests = []struct {
		args []string
		conf FlagConfig
	}{
		{[]string{},
			FlagConfig{port: 0, args: []string{}}},

		{[]string{"--port", "5000"},
			FlagConfig{port: 5000, args: []string{}}},

		{[]string{"--port", "5000", "6000"},
			FlagConfig{port: 5000, args: []string{"6000"}}},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.args, " "), func(t *testing.T) {
			conf, output, err := parseFlags("prog", tt.args)
			if err != nil {
				t.Errorf("err got %v, want nil", err)
			}
			if output != "" {
				t.Errorf("output got %q, want empty", output)
			}
			if !reflect.DeepEqual(*conf, tt.conf) {
				t.Errorf("conf got %+v, want %+v", *conf, tt.conf)
			}
		})
	}
}

func TestParseFlagsError(t *testing.T) {
	var tests = []struct {
		args   []string
		errstr string
	}{
		{[]string{"--foo"}, "flag provided but not defined"},
		{[]string{"--port"}, "flag needs an argument"},
		{[]string{"--port", "test"}, "invalid value"},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.args, " "), func(t *testing.T) {
			conf, output, err := parseFlags("prog", tt.args)
			if conf != nil {
				t.Errorf("conf got %v, want nil", conf)
			}
			if !strings.Contains(err.Error(), tt.errstr) {
				t.Errorf("err got %q, want to find %q", err.Error(), tt.errstr)
			}
			if !strings.Contains(output, "Usage of prog") {
				t.Errorf("output got %q", output)
			}
		})
	}
}
