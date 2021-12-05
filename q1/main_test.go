package main

import (
	"reflect"
	"testing"
)

func TestIncreasesCount(t *testing.T) {
	example := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	got := IncreasesCount(example)
	if got != 7 {
		t.Errorf("IncreasesCount(example) = %d; want 7", got)
	}
}

func TestWindowSums(t *testing.T) {
	example := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	expected := []int{607, 618, 618, 617, 647, 716, 769, 792}
	got := WindowSums(example, 3)
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("WindowSums(example) = %v; want %v", got, expected)
	}
}
