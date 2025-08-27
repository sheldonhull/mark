package renderer

import (
	"testing"
)

func TestMkDocsAdmonitionType_String(t *testing.T) {
	tests := []struct {
		admonitionType MkDocsAdmonitionType
		expected       string
	}{
		{AInfo, "info"},
		{ANote, "note"},
		{AWarn, "warning"},
		{ATip, "tip"},
		{AAbstract, "abstract"},
		{ASuccess, "success"},
		{AQuestion, "question"},
		{AFailure, "failure"},
		{ADanger, "danger"},
		{ABug, "bug"},
		{AExample, "example"},
		{AQuote, "quote"},
		{ANone, "none"},
	}

	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			result := test.admonitionType.String()
			if result != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, result)
			}
		})
	}
}

func TestMkDocsAdmonitionType_ConfluenceMacroName(t *testing.T) {
	tests := []struct {
		admonitionType MkDocsAdmonitionType
		expected       string
		description    string
	}{
		{AInfo, "info", "info should map to info"},
		{AAbstract, "info", "abstract should map to info"},
		{AQuestion, "info", "question should map to info"},
		{ANote, "note", "note should map to note"},
		{AQuote, "note", "quote should map to note"},
		{AWarn, "warning", "warning should map to warning"},
		{AFailure, "warning", "failure should map to warning"},
		{ADanger, "warning", "danger should map to warning"},
		{ABug, "warning", "bug should map to warning"},
		{ATip, "tip", "tip should map to tip"},
		{ASuccess, "tip", "success should map to tip"},
		{AExample, "tip", "example should map to tip"},
		{ANone, "note", "none should fallback to note"},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result := test.admonitionType.ConfluenceMacroName()
			if result != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, result)
			}
		})
	}
}