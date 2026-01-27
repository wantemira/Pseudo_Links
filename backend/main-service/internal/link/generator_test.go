package link

import "testing"

func TestGeneratePseudoLink(t *testing.T) {
	if pseudoLink := generatePseudoLink(); len(pseudoLink) <= 0 {
		t.Errorf("TestGeneratePseudoLink: error with generationg link %s", pseudoLink)
	}
}
