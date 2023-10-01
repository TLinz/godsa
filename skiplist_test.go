package godsa

import (
	"math/rand"
	"testing"
	"time"
)

func TestRandomInsertAndSearch(t *testing.T) {
	sl := NewSkipList(32, 0.5)

	rand.Seed(time.Now().UnixNano())

	keyValuePairs := make(map[int]int)
	for i := 0; i < 100; i++ {
		key := rand.Intn(1000)
		value := rand.Intn(10000)
		keyValuePairs[key] = value
		err := sl.Insert(key, value)
		if err != nil {
			t.Errorf("Insertion failed for key %d: %s", key, err.Error())
		}
	}

	for key, expectedValue := range keyValuePairs {
		result, err := sl.Search(key)
		if err != nil {
			t.Errorf("Search failed for key %d: %s", key, err.Error())
		}
		if result.value != expectedValue {
			t.Errorf("Incorrect value for key %d. Expected: %d, Actual: %d", key, expectedValue, result.value)
		}
	}
}
