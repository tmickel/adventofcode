package main

import (
	"log"
	"reflect"
	"testing"
)

func TestGamma(t *testing.T) {
	example := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	got := Gamma(example)
	if got != 22 {
		t.Errorf("Gamma(example) = %d; want 22", got)
	}
}

func TestEpsilon(t *testing.T) {
	example := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	got := Epsilon(example)
	if got != 9 {
		t.Errorf("Epsilon(example) = %d; want 9", got)
	}
}

func TestOxygen(t *testing.T) {
	example := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	got := Oxygen(example)
	if got != 23 {
		t.Errorf("Oxygen(example) = %d; want 23", got)
	}
}

func TestCO2Scrubber(t *testing.T) {
	example := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	got := CO2Scrubber(example)
	if got != 10 {
		t.Errorf("CO2Scrubber(example) = %d; want 10", got)
	}
}

func TestMostCommonBitZero(t *testing.T) {
	example := []int{1, 0, 0, 0, 1, 0, 1}
	got := MostCommonBit(example, 1)
	if got != 0 {
		t.Errorf("MostCommonBit(example) = %d; want 0", got)
	}
}

func TestMostCommonBitOne(t *testing.T) {
	example := []int{1, 1, 1}
	got := MostCommonBit(example, 0)
	if got != 1 {
		t.Errorf("MostCommonBit(example) = %d; want 1", got)
	}
}

func TestMostCommonBitTieOne(t *testing.T) {
	example := []int{1, 0, 1, 0}
	got := MostCommonBit(example, 1)
	if got != 1 {
		t.Errorf("MostCommonBit(example) = %d; want 1", got)
	}
}

func TestMostCommonBitTieZero(t *testing.T) {
	example := []int{1, 0, 1, 0}
	got := MostCommonBit(example, 0)
	if got != 0 {
		t.Errorf("MostCommonBit(example) = %d; want 0", got)
	}
}

func TestBitsFromPosition(t *testing.T) {
	example := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	expected := []int{0, 1, 1, 1, 1, 0, 0, 1, 1, 1, 0, 0}
	got := BitsFromPosition(example, 0)
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("BitsFromPosition(example) = %v; want %v", got, expected)
	}
}

func TestBugFix(t *testing.T) {
	example := []string{"11110", "10110", "10111", "10101", "11100", "10000", "11001"}
	log.Println(BitsFromPosition(example, 1))
	got := MostCommonBit(BitsFromPosition(example, 1), 1)
	expected := 0
	if got != expected {
		t.Errorf("MostCommonBit(BitsFromPosition(example, 1), 1) = %v; want %v", got, expected)
	}
}
