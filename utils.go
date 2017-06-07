package rajirec

import (
	"log"
	"net/http"
	"io/ioutil"
	"encoding/xml"
)

func AppendAllIfMissing(slice []int, i []int) []int {
	for _, ele := range i {
		slice = AppendIfMissing(slice, ele)
	}
	return slice
}

func AppendIfMissing(slice []int, i int) []int {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

type Map struct {
	children map[string]*Map
	value interface{}
}

func NewMap() *Map {
	return &Map{children: make(map[string]*Map), value: nil}
}

func (m *Map) get(key string) *Map {
	val, ok := m.children[key]
	if ok {
		if val == nil {
			return nil
		} else {
			return val
		}
	} else {
		m.children[key] = NewMap()
		return m.children[key]
	}
}

func (m *Map) Get(keys ...string) *Map {
	m1 := m
	for _, key := range keys {
		m1 = m1.get(key)
	}
	return m1
}

func (m *Map) Set(value interface{}, keys ...string) {
	m1 := m
	for _, key := range keys {
		m1 = m1.Get(key)
	}
	m1.value = value
}

func (m *Map) GetValue() interface{} {
	return m.value
}

func FetchXML(url string, v interface{}) {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Fatal("Failed to fetch XML data")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err := xml.Unmarshal(body, v); err != nil {
		log.Fatal(err)
	}
}
