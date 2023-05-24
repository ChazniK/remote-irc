package main

import "testing"

func TestProcessNickCommand(t *testing.T) {
	tests := []struct {
		message  message
		expected string
	}{
		{
			message:  message{command: "NICK", parameters: []string{"Alice"}},
			expected: "Welcome to the Internet Relay Network<nick>!<user>@<host> Alice",
		},
		{
			message:  message{command: "NICK", parameters: []string{"Jon"}},
			expected: "Jon :Nickname is already in use",
		},
	}

	for _, test := range tests {
		result := processNickCommand(test.message)
		if result != test.expected {
			t.Errorf("Expected %q, but got %q", test.expected, result)
		}
	}
}
