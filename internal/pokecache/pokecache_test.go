package pokecache

import (
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	cases := []struct {
		name string
	}{
		{
			name: "New cache",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			interval := time.Millisecond * 10
			cache := NewCache(interval)

			if cache == nil {
				t.Fatalf("NewCache() returned cache of nil")
			}
		})
	}
}

func TestAddCache(t *testing.T) {
	cases := []struct {
		name  string
		key   string
		value []byte
	}{
		{
			name:  "Add cache key1",
			key:   "key1",
			value: []byte("value1"),
		},
		{
			name:  "Add cache key2",
			key:   "key2",
			value: []byte("value2"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			interval := time.Millisecond * 10
			cache := NewCache(interval)

			cache.Add(c.key, c.value)

			value, ok := cache.Get(c.key)
			if !ok {
				t.Fatalf("Add(%q) failed: key not found", c.key)
			}

			if string(value) != string(c.value) {
				t.Fatalf("Add(%q) returned %q, expected %q", c.key, value, c.value)
			}
		})
	}
}

func TestGetCache(t *testing.T) {
	cases := []struct {
		name        string
		setupKey    string
		setupValue  []byte
		getKey      string
		expectOk    bool
		expectValue []byte
	}{
		{
			name:        "Get existing key",
			setupKey:    "key1",
			setupValue:  []byte("value1"),
			getKey:      "key1",
			expectOk:    true,
			expectValue: []byte("value1"),
		},
		{
			name:     "Get missing key",
			getKey:   "missing",
			expectOk: false,
		},
		{
			name:       "Get different key",
			setupKey:   "key1",
			setupValue: []byte("value1"),
			getKey:     "key2",
			expectOk:   false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			interval := time.Millisecond * 10
			cache := NewCache(interval)

			if c.setupKey != "" {
				cache.Add(c.setupKey, c.setupValue)
			}

			value, ok := cache.Get(c.getKey)

			if ok != c.expectOk {
				t.Fatalf("Get(%q) does nam match expected %v", c.getKey, c.expectOk)
			}

			if !c.expectOk {
				return
			}

			if string(value) != string(c.expectValue) {
				t.Fatalf("Get(%q) returned %q, expected %q", c.getKey, value, c.expectValue)
			}
		})
	}
}

func TestReap(t *testing.T) {
	cases := []struct {
		name          string
		setupEntries  map[string][]byte
		sleepDuration time.Duration
		reapInterval  time.Duration
		expectPresent []string
		expectMissing []string
	}{
		{
			name: "Removes expired entries",
			setupEntries: map[string][]byte{
				"key1": []byte("value1"),
				"key2": []byte("value2"),
			},
			sleepDuration: time.Millisecond * 20,
			reapInterval:  time.Millisecond * 10,
			expectMissing: []string{"key1", "key2"},
		},
		{
			name: "Keeps unexpired entries",
			setupEntries: map[string][]byte{
				"key1": []byte("value1"),
			},
			sleepDuration: time.Millisecond * 5,
			reapInterval:  time.Millisecond * 10,
			expectPresent: []string{"key1"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cache := NewCache(time.Hour)

			for k, v := range c.setupEntries {
				cache.Add(k, v)
			}

			time.Sleep(c.sleepDuration)

			cache.Reap(c.reapInterval)

			for _, key := range c.expectPresent {
				if _, ok := cache.Get(key); !ok {
					t.Fatalf("expected key %q to be present after reap", key)
				}
			}

			for _, key := range c.expectMissing {
				if _, ok := cache.Get(key); ok {
					t.Fatalf("expected key %q to be removed after reap", key)
				}
			}
		})
	}
}
