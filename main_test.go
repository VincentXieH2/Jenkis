package main

import (
        "testing"
)

func TestAdd(t *testing.T) {
        result := Add(1, 5)
        if result != 4 {
                t.Errorf("should be 4 but %d", result)
        }
}
