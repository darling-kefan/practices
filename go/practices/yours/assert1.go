package yours

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	assert := assert.New(t)

	// assert equality
	assert.Equal(t, 1231, 123, "they should be equal.")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

	var object = struct{
		Value string
	}{
		Value: "hello world",
	}
	
	// assert for nil (good for errors)
	assert.Nil(t, object)

	// assert for nil (good for errors)
	if assert.NotNil(t, object) {
		// now we known that object isn't nil, we are safe to make
		// further assertions without causing any errors.
		assert.Equal(t, "something", object.Value)
	}
}
