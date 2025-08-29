package parser

import (
	"testing"
)

func TestEnhancedAdmonitionParser_Trigger(t *testing.T) {
	p := NewEnhancedAdmonitionParser()
	triggers := p.Trigger()
	
	// Should support both ! and ? characters
	if len(triggers) != 2 {
		t.Errorf("Expected 2 trigger characters, got %d", len(triggers))
	}
	
	if triggers[0] != '!' {
		t.Errorf("Expected first trigger to be '!', got %c", triggers[0])
	}
	
	if triggers[1] != '?' {
		t.Errorf("Expected second trigger to be '?', got %c", triggers[1])
	}
}