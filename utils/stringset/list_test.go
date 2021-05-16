package stringset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntsToList(t *testing.T) {
	cases := []struct {
		delim    byte
		expected string
		data     []int
	}{
		{',', "", []int{}},
		{',', "100", []int{100}},
		{',', "1,2,9,10", []int{1, 2, 9, 10}},
		{',', "98,-321,1230,-98765,0", []int{98, -321, 1230, -98765, 0}},
		{' ', "98 -321 1230 -98765 0", []int{98, -321, 1230, -98765, 0}},
	}
	for _, c := range cases {
		actual := ListFromIntSlice(c.data, c.delim)
		assert.EqualValues(t, c.expected, actual, "int slice to List")
	}
}

func TestSliceToList(t *testing.T) {
	cases := []struct {
		delim    byte
		expected string
		data     interface{}
	}{
		{',', "", []float64{}},
		{',', "", nil},
		{',', "sopas", []string{"sopas"}},
		{',', "1.1,2.4,9.89,10.45", []float64{1.1, 2.4, 9.89, 10.45}},
		{',', "true,true,false,true,false", []bool{true, true, false, true, false}},
		{' ', "97 98 99 100", []rune{'a', 'b', 'c', 'd'}},
	}
	for _, c := range cases {
		actual := ListFromSlice(c.data, c.delim)
		assert.EqualValues(t, c.expected, actual, "int slice to List")
	}
}

func TestToSlice(t *testing.T) {
	cases := []struct {
		delim    byte
		data     string
		expected []string
	}{
		{',', "", []string{}},
		{',', "sopas", []string{"sopas"}},
		{',', "1.1,2.4,9.89,10.45", []string{"1.1", "2.4", "9.89", "10.45"}},
		{',', "true,true,false,true,false", []string{"true", "true", "false", "true", "false"}},
		{' ', "a b c d", []string{"a", "b", "c", "d"}},
	}
	for _, c := range cases {
		actual := ListToSlice(c.data, c.delim)
		assert.EqualValues(t, c.expected, actual, "List to slice")
	}
}
