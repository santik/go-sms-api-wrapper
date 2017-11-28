package main

import "testing"

func TestMessageMaxLength(t *testing.T) {
	m := Message{}
	if m.max_message_length() != 160 {
		t.Error("Expected 60, got ", m.max_message_length())
	}
}