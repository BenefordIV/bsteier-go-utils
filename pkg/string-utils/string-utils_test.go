package string_utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEqualsIgnoreCase(t *testing.T) {
	target := "target"
	v := "TaRget"

	e := EqualsIgnoreCase(target, v)

	assert.True(t, e)
}

func TestContainsIgnoreCase(t *testing.T) {
	target := "target"
	v := "tar"

	e := ContainsIgnoreCase(target, v)

	assert.True(t, e)
}
