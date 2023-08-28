package dynamic_utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionalStructFieldsToMap_WithStruct(t *testing.T) {
	type ts struct {
		Name *string
		Age  *int
	}

	var rm map[string]any
	rm = OptionalStructFieldsToMap(ts{})
	assert.Equal(t, 0, len(rm))

	name := "name"
	rm = OptionalStructFieldsToMap(ts{Name: &name})
	assert.Equal(t, 1, len(rm))
	assert.Equal(t, name, rm["Name"])

	age := 10
	rm = OptionalStructFieldsToMap(ts{Name: &name, Age: &age})
	assert.Equal(t, 2, len(rm))
	assert.Equal(t, name, rm["Name"])
	assert.Equal(t, age, rm["Age"])
}

func TestOptionalStructFieldsToMap_WithStructPointer(t *testing.T) {
	type ts struct {
		Name *string
		Age  *int
	}

	var rm map[string]any
	rm = OptionalStructFieldsToMap(&ts{})
	assert.Equal(t, 0, len(rm))

	name := "name"
	rm = OptionalStructFieldsToMap(&ts{Name: &name})
	assert.Equal(t, 1, len(rm))
	assert.Equal(t, name, rm["Name"])

	age := 10
	rm = OptionalStructFieldsToMap(&ts{Name: &name, Age: &age})
	assert.Equal(t, 2, len(rm))
	assert.Equal(t, name, rm["Name"])
	assert.Equal(t, age, rm["Age"])
}
