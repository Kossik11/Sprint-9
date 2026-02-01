package main

import "testing"

func TestGenerateRandomElements_SizeZero(t *testing.T) {
  data := generateRandomElements(0)
  if data == nil {
    t.Fatalf("ожидали непустой слайс, получили nil")
  }
  if len(data) != 0 {
    t.Fatalf("ожидали длину 0, получили %d", len(data))
  }
}

func TestGenerateRandomElements_NegativeSize(t *testing.T) {
  data := generateRandomElements(-10)
  if data == nil {
    t.Fatalf("ожидали непустой слайс, получили nil")
  }
  if len(data) != 0 {
    t.Fatalf("ожидали длину 0, получили %d", len(data))
  }
}

func TestGenerateRandomElements_PositiveSize(t *testing.T) {
  size := 1000
  data := generateRandomElements(size)

  if data == nil {
    t.Fatalf("ожидали непустой слайс, получили nil")
  }
  if len(data) != size {
    t.Fatalf("ожидали длину %d, получили %d", size, len(data))
  }

  for i, v := range data {
    if v <= 0 {
      t.Fatalf("ожидали положительное число в data[%d], получили %d", i, v)
    }
  }
}

func TestMaximum_EmptySlice(t *testing.T) {
  if got := maximum([]int{}); got != 0 {
    t.Fatalf("ожидали 0 для пустого слайса, получили %d", got)
  }
}

func TestMaximum_NilSlice(t *testing.T) {
  var data []int
  if got := maximum(data); got != 0 {
    t.Fatalf("ожидали 0 для nil-слайса, получили %d", got)
  }
}

func TestMaximum_OneElement(t *testing.T) {
  if got := maximum([]int{7}); got != 7 {
    t.Fatalf("ожидали 7, получили %d", got)
  }
}

func TestMaximum_ManyElements(t *testing.T) {
  data := []int{3, 1, 9, 2, 9, 5}
  got := maximum(data)
  want := 9
  if got != want {
    t.Fatalf("ожидали %d, получили %d", want, got)
  }
}

func TestMaximum_AllEqual(t *testing.T) {
  data := []int{4, 4, 4}
  got := maximum(data)
  want := 4
  if got != want {
    t.Fatalf("ожидали %d, получили %d", want, got)
  }
}

func TestMaxChunks_Empty(t *testing.T) {
  if got := maxChunks([]int{}); got != 0 {
    t.Fatalf("ожидали 0 для пустого слайса, получили %d", got)
  }
}