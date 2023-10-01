package godsa

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

func TestSkipListOperations(t *testing.T) {
	assert := assert.New(t)
	maxLevel := 16
	probability := 0.5
	sl := NewSkipList(maxLevel, probability)

	keys := []int{10, 5, 30, 15, 20}
	values := []int{100, 50, 300, 150, 200}

	for i, key := range keys {
		err := sl.Insert(key, values[i])
		assert.NoError(err, "Insert should not return an error")
	}

	for i, key := range keys {
		node, err := sl.Search(key)
		assert.NoError(err, "Search should not return an error")
		assert.NotNil(node, "Search result should not be nil")
		assert.Equal(node.value, values[i], "Search result value should match")
	}

	err := sl.Delete(15)
	assert.NoError(err, "Delete should not return an error")
	_, err = sl.Search(15)
	assert.Error(err, "Search for deleted element should return an error")

	searchKeys := []int{10, 5, 30, 20}
	for _, key := range searchKeys {
		_, err := sl.Search(key)
		assert.NoError(err, "Search should not return an error")
	}

	err = sl.Delete(25)
	assert.Error(err, "Delete for non-existent element should return an error")
}
