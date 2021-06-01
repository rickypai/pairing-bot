package main

import (
	"fmt"
	"testing"
)

var daysList = map[string]struct{}{
	"monday":    {},
	"tuesday":   {},
	"wednesday": {},
	"thursday":  {},
	"friday":    {},
	"saturday":  {},
	"sunday":    {},
}

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
	{"Help", "help", nil, false},
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

var tableWithArgs = []struct {
	inputStr   string
	wantedCmd  string
	wantedArgs []string
	expectErr  bool
}{
	{"schedule monday", "schedule", []string{"monday"}, false},
	{"schedule monday friday", "schedule", []string{"monday", "friday"}, false},
	{"schedule", "help", nil, true},
	{"skip tomorrow", "skip", []string{"tomorrow"}, false},
	{"skip whenever", "help", nil, true},
	{"skip", "help", nil, true},
	{"unskip tomorrow", "unskip", []string{"tomorrow"}, false},
	{"unskip today", "help", nil, true},
	{"unskip", "help", nil, true},
}

func TestParseCmdWithArgs(t *testing.T) {
	for _, tt := range tableWithArgs {
		testname := fmt.Sprintf("\nTesting command %v, expecting %v, %v, and %v\n", tt.inputStr, tt.wantedCmd, tt.wantedArgs, tt.expectErr)
		t.Run(testname, func(t *testing.T) {
			gotCmd, gotArgs, gotErr := parseCmd(tt.inputStr)
			if gotCmd != tt.wantedCmd || len(gotArgs) != len(tt.wantedArgs) {
				t.Errorf("got %v, %v, wanted %v, %v\n", gotCmd, gotArgs, tt.wantedCmd, tt.wantedArgs)
			}

			switch gotCmd {
			case "schedule":
				for i, arg := range gotArgs {
					if _, ok := daysList[arg]; !ok {
						t.Errorf("Wrong argument %v for command %v\n", gotArgs[i], gotCmd)
					}
				}
			case "skip":
				if gotArgs[0] != "tomorrow" {
					t.Errorf("Wrong argument %v for command %v\n", gotArgs[0], gotCmd)
				}
			case "unskip":
				if gotArgs[0] != "tomorrow" {
					t.Errorf("Wrong argument %v for command %v\n", gotArgs[0], gotCmd)
				}
			default:
				if gotCmd != "help" {
					t.Errorf("unknown command %v\n", gotCmd)
				}
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

var tableInvalid = []struct {
	inputStr   string
	wantedCmd  string
	wantedArgs []string
	expectErr  bool
}{
	{"scheduleing monday", "help", nil, true},
	{"mooh", "help", nil, true},
}

func TestInvalidCommands(t *testing.T) {
	for _, tt := range tableInvalid {
		testname := fmt.Sprintf("\nTesting command %v, expecting %v, %v, and %v\n", tt.inputStr, tt.wantedCmd, tt.wantedArgs, tt.expectErr)
		t.Run(testname, func(t *testing.T) {
			gotCmd, gotArgs, gotErr := parseCmd(tt.inputStr)
			if gotCmd != tt.wantedCmd {
				t.Errorf("got %v, %v, wanted %v, %v\n", gotCmd, gotArgs, tt.wantedCmd, tt.wantedArgs)
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
