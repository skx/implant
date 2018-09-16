package main

import (
	"os"
	"testing"
)

// TestParseArguments tests our argument-parsing.  (Trivial/NOP)
func TestParseArguments(t *testing.T) {

	ConfigOptions.Input = nil
	ConfigOptions.Exclude = ""
	ConfigOptions.Output = ""
	ConfigOptions.Package = "main"
	ConfigOptions.Verbose = false
	ConfigOptions.Format = false
	ConfigOptions.Version = false

	os.Args = []string{"cmd", "-input=foo", "-input=bar", "-package=leather", "-verbose"}

	parseArguments()

	if len(ConfigOptions.Input) != 2 {
		t.Errorf("Expected to input directories, got %d - %v", len(ConfigOptions.Input), ConfigOptions.Input)
	}
	if ConfigOptions.Verbose == false {
		t.Errorf("Parsing -verbose flag failed")
	}
	if ConfigOptions.Package != "leather" {
		t.Errorf("Parsing -package flag failed")
	}
}
