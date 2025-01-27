package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Second * 1)

	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	cache := NewCache(time.Second * 1)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "",
			inputVal: []byte("val3"),
		},
	}

	for _, cas := range cases {
		cache.Add(cas.inputKey, cas.inputVal)
		actual, ok := cache.Get(cas.inputKey)
		if !ok {
			t.Errorf("%s not found", cas.inputKey)
			continue
		}

		if string(actual) != string(cas.inputVal) {
			t.Errorf("%s not found %s", string(actual), cas.inputVal)
			continue
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10

	cache := NewCache(interval)

	cache.Add("key1", []byte("val1"))
	time.Sleep(interval + time.Millisecond*5)
	_, ok := cache.Get("key1")
	if ok {
		t.Errorf("key1 should have been reaped")
	}
}

func TestReapFails(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	cache.Add("key2", []byte("val2"))
	time.Sleep(time.Millisecond)
	_, ok := cache.Get("key2")
	if !ok {
		t.Errorf("key2 should have been found")
	}
}
