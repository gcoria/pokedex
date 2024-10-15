package pokecache

import "testing"

func TestCreateCache(t *testing.T) {
	cache := NewCache()

	if cache.cache == nil {
		t.Error("cache is nil")
	}
}
