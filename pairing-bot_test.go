package main

import (
	"fmt"
	"testing"
)

// due to the difficulties around comparing error return values (and because we don't want to compare error messages),
// the struct contains expectErr to indicate whether an error is expected, instead of an actual error value
var tableNoArgs = []struct {
	inputStr   string
	wantedCmd  string
	wantedArgs []string
	expectErr  bool
}{
	{"subscribe", "subscribe", nil, false},
	{"subscribe mon", "help", nil, true},
	{"unsubscribe", "unsubscribe", nil, false},
	{"unsubscribe tuesday", "help", nil, true},
	{"help", "help", nil, false},
	{"help me", "help", nil, true},
	{"status", "status", nil, false},
	{"status me", "help", nil, true},
	{"count", "count", nil, false},
	{"count me", "help", nil, true},
}

func TestParseCmdNoArgs(t *testing.T) {
	for _, tt := range tableNoArgs {
		testname := fmt.Sprintf("\nTesting command %v, expecting %v, %v, and %v\n", tt.inputStr, tt.wantedCmd, tt.wantedArgs, tt.expectErr)
		t.Run(testname, func(t *testing.T) {
			gotCmd, gotArgs, gotErr := parseCmd(tt.inputStr)
			if gotCmd != tt.wantedCmd {
				t.Errorf("got %v, %v\n", gotCmd, gotArgs)
			}

			_, ok := gotErr.(*parsingErr)

			if tt.expectErr && !ok {
				t.Errorf("Expected parsingErr but didn't get one\n")
			} else if !tt.expectErr && ok {
				t.Errorf("Got unexpected parsingError\n")
			}
		})
	}
}
