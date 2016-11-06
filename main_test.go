package main_test

import (
    "testing"

    . "github.com/brianjarita/go_test_idea"
)


func TestAdd(t *testing.T) {
  t.Run("Adding 1", func(t *testing.T) {
    got := Add(1,1)
    if got != 2 {
      t.Errorf("got %d, want 2", got)
    }
  })
  t.Run("Add 2", func(t *testing.T) {
    got := Add(1, 2)
    if got != 3 {
      t.Errorf("got %d, want 3", got)
    }
  })
}

func TestSub(t *testing.T) {
  t.Run("Subtracting 1", func(t *testing.T) {
    got := Subtract(5,1)
    if got != 4 {
      t.Errorf("got %d, want 4", got)
    }
  })
  t.Run("Subtracting 2", func(t *testing.T) {
    got := Subtract(5, 2)
    if got != 3 {
      t.Errorf("got %d, want 3", got)
    }
  })
}
