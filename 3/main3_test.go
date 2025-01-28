package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	m := NewStringIntMap()
	m.Add("one", 1)
	m.Add("two", 2)

	if value, exists := m.Get("one"); !exists || value != 1 {
		t.Errorf("Expected value 1 for key 'one', got %v", value)
	}
	if value, exists := m.Get("two"); !exists || value != 2 {
		t.Errorf("Expected value 2 for key 'two', got %v", value)
	}
}

func TestRemove(t *testing.T) {
	m := NewStringIntMap()
	m.Add("one", 1)
	m.Add("two", 2)

	m.Remove("one")
	if exists := m.Exists("one"); exists {
		t.Error("Key 'one' should not exist after removal")
	}

	if exists := m.Exists("two"); !exists {
		t.Error("Key 'two' should exist after removal of 'one'")
	}
}

func TestCopy(t *testing.T) {
	m := NewStringIntMap()
	m.Add("one", 1)
	m.Add("two", 2)

	copiedMap := m.Copy()
	if len(copiedMap) != 2 {
		t.Errorf("Expected copied map length 2, got %v", len(copiedMap))
	}

	if value, exists := copiedMap["one"]; !exists || value != 1 {
		t.Errorf("Expected value 1 for key 'one' in copied map, got %v", value)
	}
}

func TestExists(t *testing.T) {
	m := NewStringIntMap()
	m.Add("one", 1)

	if exists := m.Exists("one"); !exists {
		t.Error("Key 'one' should exist")
	}

	if exists := m.Exists("two"); exists {
		t.Error("Key 'two' should not exist")
	}
}

func TestGet(t *testing.T) {
	m := NewStringIntMap()
	m.Add("one", 1)

	if value, exists := m.Get("one"); !exists || value != 1 {
		t.Errorf("Expected value 1 for key 'one', got %v", value)
	}

	if _, exists := m.Get("two"); exists {
		t.Error("Key 'two' should not exist")
	}
}
