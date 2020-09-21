package mu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomString(t *testing.T) {
	SetCharset(AlphaNumeric)
	str := RandomString(10)
	assert.Equal(t, len(str), 10)
}
