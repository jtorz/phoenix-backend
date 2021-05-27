package fndmodel

import (
	"fmt"
	"testing"
)

func TestNavElementsAdd(t *testing.T) {
	Z1 := NavElement{ID: "Z.1", Order: 10, ParentID: "Z"}
	X := NavElement{ID: "X", Order: 10}
	X1 := NavElement{ID: "X.1", ParentID: "X", Order: 10}
	X11 := NavElement{ID: "X.1.1", ParentID: "X.1", Order: 10}
	X12 := NavElement{ID: "X.1.2", ParentID: "X.1", Order: 20}
	X13 := NavElement{ID: "X.1.3", ParentID: "X.1", Order: 30}
	X14 := NavElement{ID: "X.1.4", ParentID: "X.1", Order: 40}
	X2 := NavElement{ID: "X.2", ParentID: "X", Order: 20}
	X3 := NavElement{ID: "X.3", ParentID: "X", Order: 30}
	X4 := NavElement{ID: "X.4", ParentID: "X", Order: 40}
	Y := NavElement{ID: "Y", Order: 20}
	Y1 := NavElement{ID: "Y.1", ParentID: "Y", Order: 10}
	Y2 := NavElement{ID: "Y.2", ParentID: "Y", Order: 20}
	Y3 := NavElement{ID: "Y.3", ParentID: "Y", Order: 30}
	Y31 := NavElement{ID: "Y.3.1", ParentID: "Y.3", Order: 10}
	Y32 := NavElement{ID: "Y.3.2", ParentID: "Y.3", Order: 20}
	Y33 := NavElement{ID: "Y.3.3", ParentID: "Y.3", Order: 30}
	Y34 := NavElement{ID: "Y.3.4", ParentID: "Y.3", Order: 40}
	Y4 := NavElement{ID: "Y.4", ParentID: "Y", Order: 40}
	records := Navigator{
		Z1,
		X,
		X2,
		X4,
		X3,
		X1,
		Y1,
		Y3,
		Y2,
		Y4,
		X11,
		X12,
		X13,
		X14,
		Y31,
		Y33,
		Y32,
		Y34,
		Y,
	}

	tree := records.Tree()
	X1.Children = Navigator{X11, X12, X13, X14}
	X.Children = Navigator{X1, X2, X3, X4}
	Y3.Children = Navigator{Y31, Y32, Y33, Y34}
	Y.Children = Navigator{Y1, Y2, Y3, Y4}

	//expectedTree := Navigator{X, Y, Z1}
	//assert.EqualValues(t, expectedTree, tree)

	tree.print(0)
}

func TestAppend(t *testing.T) {
	a := []float64{0, 1, 2, 3, 4, 5, 6, 7}

	val := 8.0
	var i int
	for i = 0; i < len(a); i++ {
		if val < a[i] {
			break
		}
	}

	first := make([]float64, i)
	copy(first, a[:i])
	second := a[i:]
	fmt.Println(append(append(first, val), second...))
}
