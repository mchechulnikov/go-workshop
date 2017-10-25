package hashMap

import (
	"fmt"
	"go-workshop/data-structures/linkedListtures/linkedList"
	"hash/fnv"
)

func New() *HashMap {
	return &HashMap{support: make([]*linkedList.LinkedList, 10)}
}

type HashMap struct {
	support []*linkedList.LinkedList
	length  int
}

func (m *HashMap) Length() int {
	return m.length
}

func (m *HashMap) Get(key string) (interface{}, error) {
	bucketIndex := m.calcIndex(key)
	if m.support[bucketIndex] == nil {
		return nil, fmt.Errorf("No value presented for key '%s'", key)
	}
	bucketElement := m.support[bucketIndex].First
	for {
		kv := bucketElement.GetValue().(*kvPair)
		if kv.key == key {
			return kv.value, nil
		}
		bucketElement = bucketElement.GetNext()
		if bucketElement == nil {
			return nil, fmt.Errorf("No value presented for key '%s'", key)
		}
	}
}

func (m *HashMap) Add(key string, value interface{}) error {
	if v, _ := m.Get(key); v != nil {
		return fmt.Errorf("Key '%s' is already presented in hash map", key)
	}
	index := m.calcIndex(key)
	if m.support[index] == nil {
		m.support[index] = linkedList.New()
	}
	m.support[index].Add(&kvPair{key, value})
	m.length++
	return nil
}

func (m *HashMap) Delete(key string) error {
	bucketIndex := m.calcIndex(key)
	if m.support[bucketIndex] == nil {
		return fmt.Errorf("No value presented for key '%s'", key)
	}
	bucketElement := m.support[bucketIndex].First
	index := 0
	for {
		kv := bucketElement.GetValue().(*kvPair)
		if kv.key == key {
			m.support[bucketIndex].Delete(index)
			m.length--
			return nil
		}
		bucketElement = bucketElement.GetNext()
		if bucketElement == nil {
			return fmt.Errorf("No value presented for key '%s'", key)
		}
		index++
	}
}

func (m *HashMap) calcIndex(key string) int {
	if key == "" {
		return 0
	}

	hashCode := calcHashCode(key)
	return (hashCode ^ (hashCode >> 16)) & (len(m.support) - 1)
}

func calcHashCode(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32())
}

type kvPair struct {
	key   string
	value interface{}
}
