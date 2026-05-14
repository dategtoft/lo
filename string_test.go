package lo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCapitalize(t *testing.T) {
	assert.Equal(t, "Hello", Capitalize("hello"))
	assert.Equal(t, "Hello", Capitalize("Hello"))
	assert.Equal(t, "H", Capitalize("h"))
	assert.Equal(t, "", Capitalize(""))
}

func TestCamelCase(t *testing.T) {
	assert.Equal(t, "helloWorld", CamelCase("hello_world"))
	assert.Equal(t, "helloWorld", CamelCase("hello-world"))
	assert.Equal(t, "helloWorld", CamelCase("Hello World"))
	assert.Equal(t, "helloWorld", CamelCase("HelloWorld"))
	assert.Equal(t, "", CamelCase(""))
}

func TestSnakeCase(t *testing.T) {
	assert.Equal(t, "hello_world", SnakeCase("helloWorld"))
	assert.Equal(t, "hello_world", SnakeCase("hello-world"))
	assert.Equal(t, "hello_world", SnakeCase("Hello World"))
	assert.Equal(t, "hello_world", SnakeCase("hello_world"))
	assert.Equal(t, "", SnakeCase(""))
}

func TestKebabCase(t *testing.T) {
	assert.Equal(t, "hello-world", KebabCase("helloWorld"))
	assert.Equal(t, "hello-world", KebabCase("hello_world"))
	assert.Equal(t, "hello-world", KebabCase("Hello World"))
	assert.Equal(t, "", KebabCase(""))
}

func TestPascalCase(t *testing.T) {
	assert.Equal(t, "HelloWorld", PascalCase("hello_world"))
	assert.Equal(t, "HelloWorld", PascalCase("hello-world"))
	assert.Equal(t, "HelloWorld", PascalCase("helloWorld"))
	assert.Equal(t, "", PascalCase(""))
}

func TestTruncate(t *testing.T) {
	assert.Equal(t, "hello", Truncate("hello", 10, "..."))
	assert.Equal(t, "he...", Truncate("hello world", 5, "..."))
	assert.Equal(t, "hello", Truncate("hello", 5, "..."))
	assert.Equal(t, "...", Truncate("hello world", 3, "..."))
	assert.Equal(t, "", Truncate("", 5, "..."))
}
