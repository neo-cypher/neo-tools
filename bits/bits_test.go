package bits

import (
	"fmt"
	"math"
	"testing"
)

func TestBitCount32(t *testing.T) {
	tests := []struct {
		input, expected uint32
	}{
		{0b0000_0000_0000_0000, 0},
		{0b0000_0000_0000_0001, 1},
		{0b0000_0000_0000_1111, 4},
		{0b1111_1111_1111_1111, 16},
		{math.MaxUint32, 32},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%v", test.input)
		t.Run(name, func(t *testing.T) {
			actual := BitCount32(test.input)
			if actual != test.expected {
				t.Errorf("got %v, expected %v", actual, test.expected)
			}
		})
	}
}

func TestXor(t *testing.T) {
	tests := []struct {
		left, right, expected uint32
	}{
		{0b0000_0000, 0b0000_0000, 0b0000_0000},
		{0b0000_0000, 0b0000_0001, 0b0000_0001},
		{0b0000_0001, 0b0000_0000, 0b0000_0001},
		{0b0000_0001, 0b0000_0001, 0b0000_0000},
		{0b1010_1111, 0b0101_1111, 0b1111_0000},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%v^%v", test.left, test.right)
		t.Run(name, func(t *testing.T) {
			actual := Xor(test.left, test.right)
			if actual != test.expected {
				t.Errorf("got %v, expected %v", actual, test.expected)
			}
		})
	}
}

func TestCheckBit(t *testing.T) {
	tests := []struct {
		value, bit uint32
		expected   bool
	}{
		{0b0000_0000, 0, false},
		{0b0000_0000, 1, false},
		{0b0000_0001, 0, true},
		{0b0000_0001, 1, false},
		{0b1010_1111, 7, true},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%v,%v", test.value, test.bit)
		t.Run(name, func(t *testing.T) {
			actual := CheckBit(test.value, test.bit)
			if actual != test.expected {
				t.Errorf("got %v, expected %v", actual, test.expected)
			}
		})
	}
}

func TestSetBit(t *testing.T) {
	tests := []struct {
		value, bit, expected uint32
	}{
		{0b0000_0000, 0, 0b0000_0001},
		{0b0000_0000, 1, 0b0000_0010},
		{0b0000_0001, 0, 0b0000_0001},
		{0b0000_0001, 1, 0b0000_0011},
		{0b1010_1111, 6, 0b1110_1111},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%v,%v", test.value, test.bit)
		t.Run(name, func(t *testing.T) {
			actual := SetBit(test.value, test.bit)
			if actual != test.expected {
				t.Errorf("got %v, expected %v", actual, test.expected)
			}
		})
	}
}

func TestClearBit(t *testing.T) {
	tests := []struct {
		value, bit, expected uint32
	}{
		{0b0000_0000, 0, 0b0000_0000},
		{0b0000_0000, 1, 0b0000_0000},
		{0b0000_0001, 0, 0b0000_0000},
		{0b0000_0001, 1, 0b0000_0001},
		{0b1010_1111, 7, 0b0010_1111},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%v,%v", test.value, test.bit)
		t.Run(name, func(t *testing.T) {
			actual := ClearBit(test.value, test.bit)
			if actual != test.expected {
				t.Errorf("got %v, expected %v", actual, test.expected)
			}
		})
	}
}

func TestToggleBit(t *testing.T) {
	tests := []struct {
		value, bit, expected uint32
	}{
		{0b0000_0000, 0, 0b0000_0001},
		{0b0000_0000, 1, 0b0000_0010},
		{0b0000_0001, 0, 0b0000_0000},
		{0b0000_0001, 1, 0b0000_0011},
		{0b1010_1111, 7, 0b0010_1111},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%v,%v", test.value, test.bit)
		t.Run(name, func(t *testing.T) {
			actual := ToggleBit(test.value, test.bit)
			if actual != test.expected {
				t.Errorf("got %v, expected %v", actual, test.expected)
			}
		})
	}
}
