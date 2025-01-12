package string_utils

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestContainsIgnoreCase(t *testing.T) {
	target := "target"
	v := "taR"

	e := ContainsIgnoreCase(target, v)
	e2 := strings.Contains(target, v)

	assert.True(t, e)
	assert.False(t, e2)
}
