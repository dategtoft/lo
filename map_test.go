package lo

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	keys := Keys(m)
	sort.Strings(keys)
	assert.Equal(t, []string{"a", "b", "c"}, keys)
}

func TestValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	vals := Values(m)
	sort.Ints(vals)
	assert.Equal(t, []int{1, 2, 3}, vals)
}

func TestEntries(t *testing.T) {
	m := map[string]int{"a": 1}
	entries := Entries(m)
	assert.Len(t, entries, 1)
	assert.Equal(t, "a", entries[0].Key)
	assert.Equal(t, 1, entries[0].Value)
}

func TestFromEntries(t *testing.T) {
	entries := []Entry[string, int]{{Key: "a", Value: 1}, {Key: "b", Value: 2}}
	result := FromEntries(entries)
	assert.Equal(t, map[string]int{"a": 1, "b": 2}, result)
}

func TestMapValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	result := MapValues(m, func(v int, _ string) int { return v * 10 })
	assert.Equal(t, map[string]int{"a": 10, "b": 20}, result)
}

func TestOmitByKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	result := OmitByKeys(m, []string{"b"})
	assert.Equal(t, map[string]int{"a": 1, "c": 3}, result)
}
