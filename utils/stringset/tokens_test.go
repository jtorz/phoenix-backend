package stringset

import (
	"testing"
)

func TestA(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "Escuela", expected: "escuela"},
		{input: "Alumnos Inscritos", expected: "alumn_inscr"},
		{input: "Alumnos \"Inscritos\"", expected: "alumn_inscr"},
		{input: "Alumnos \"Inscritos\" SUPERIOR", expected: "alumn_in_s"},
		{input: "Alumnos ''INSCRITOS'' 'superior'", expected: "alumn_in_s"},
		{input: "Créditos", expected: "creditos"},
		{input: "Cré\nditos", expected: "cre_ditos"},
		{input: "Hellö\ntHëRé!", expected: "hello_there"},
		{input: "9 cats", expected: "n9_cats"},
		{input: "sexo", expected: "sexo"},
		{input: "nivel", expected: "nivel"},
		{input: "escuela", expected: "escuela"},
		{input: "tipo", expected: "tipo"},
		{input: "becados", expected: "becados"},
		{input: "monto", expected: "monto"},
		{input: "Becas", expected: "becas"},
	}
	for _, test := range tests {
		actual := NewToken(test.input)
		if actual != test.expected {
			t.Errorf("{%s} expected {%s}; but got {%s}", test.input, test.expected, actual)
		}
	}
}
