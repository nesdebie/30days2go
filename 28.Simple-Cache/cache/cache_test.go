package cache

import (
    "testing"
    "time"
)

func TestSetGet(t *testing.T) {
    c := NewCache()
    c.Set("a", 123, 100*time.Millisecond)
    v, found := c.Get("a")
    if !found || v != 123 {
        t.Error("Value not found or incorrect")
    }
}

func TestExpiration(t *testing.T) {
    c := NewCache()
    c.Set("b", "gone", 50*time.Millisecond)
    time.Sleep(100 * time.Millisecond)
    _, found := c.Get("b")
    if found {
        t.Error("Value should have expired")
    }
}

func TestDelete(t *testing.T) {
    c := NewCache()
    c.Set("c", 1, time.Second)
    c.Delete("c")
    _, found := c.Get("c")
    if found {
        t.Error("Value should be deleted")
    }
}
