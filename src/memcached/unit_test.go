package memcached

import "testing"

func TestSet(t *testing.T) {
	key := "key"
	value := "value"
	first := Get(key)
	if first != "" {
		t.Errorf("Non-empty value %v", first)
	}
	Set(key, value)
	second := Get(key)
	if second != value {
		t.Errorf("Invalid value %v", second)
	}
	Del(key)
	third := Get(key)
	if third != "" {
		t.Errorf("Non-deleted value %v", third)
	}
}
