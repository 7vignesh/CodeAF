package prose

import (
	"strings"
	"testing"
)

// TestABareNumberedReplyStillDraws guards #1072: a reply that is only a number
// and a full stop (`32.`, `1024.`) is parsed by Markdown as an ordered list
// with one empty item. The list renderer must draw the marker as the literal
// text the reader sent, never discard it as an empty list and render nothing.
func TestABareNumberedReplyStillDraws(t *testing.T) {
	for _, in := range []string{"32.", "1024."} {
		if got := strings.TrimSpace(strings.Join(Render(in, Options{Width: 80}), "\n")); got == "" {
			t.Errorf("%q rendered to nothing", in)
		}
	}
	// A genuine ordered list with a body still draws as a list.
	if got := strings.TrimSpace(strings.Join(Render("1. one", Options{Width: 80}), "\n")); got == "" {
		t.Errorf("%q rendered to nothing", "1. one")
	}
}
