package igmarkets

import (
	"testing"

	"github.com/AMekss/assert"
)

func TestNew(t *testing.T) {
	igm := New(DemoAPIURL, "", "", "", "")
	assert.False(t, igm == nil)
}
