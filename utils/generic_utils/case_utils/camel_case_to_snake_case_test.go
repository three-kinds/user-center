package case_utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCamelCaseToSnakeCase(t *testing.T) {
	s := "CamelCase"
	rs := CamelCaseToSnakeCase(s)
	assert.Equal(t, "camel_case", rs)
}
