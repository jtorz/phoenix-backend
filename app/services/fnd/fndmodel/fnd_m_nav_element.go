package fndmodel

import (
	"fmt"
	"strings"
	"time"

	"github.com/jtorz/phoenix-backend/app/shared/base"
)

// Navigator slice.
type Navigator []NavElement

// NavElement data.
type NavElement struct {
	ID          string
	Name        string
	Description string
	Icon        string
	Order       int
	URL         string
	ParentID    string
	Children    Navigator
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Status      base.Status
	base.RecordActions
	// IsAssigned to a role.
	IsAssigned bool
}

func (navs Navigator) print(lvl int) {
	for _, n := range navs {
		fmt.Printf("%s%s\n", strings.Repeat(" ", lvl), n.ID)
		if n.Children != nil {
			n.Children.print(lvl + 1)
		}
	}
}

// Tree create a tree from the navigators.
//
// This function assumes that the original slice is ordered by level.
func (navs Navigator) Tree() Navigator {
	tree := Navigator{}
	remaining := Navigator{}
	for i := range navs {
		if navs[i].ParentID == "" {
			tree = tree.appendSort(navs[i])
			continue
		}
		if ok := tree.add(navs[i]); !ok {
			remaining = remaining.appendSort(navs[i])
		}
	}

	for i := range remaining {
		tree = append(tree, navs[i])
	}
	return tree
}

// appendSort checks the Order to add the element in the correct order.
func (navs Navigator) appendSort(n NavElement) Navigator {
	l := len(navs)
	if l == 0 {
		return Navigator{n}
	}
	var i int
	for i = 0; i < l; i++ {
		if n.Order < navs[i].Order {
			break
		}
	}
	first := make(Navigator, i)
	copy(first, navs[:i])
	return append(append(first, n), navs[i:]...)
}

func (navs *Navigator) add(nav NavElement) bool {
	for i := range *navs {
		if (*navs)[i].addIfChild(nav) {
			return true
		} else if (*navs)[i].Children.add(nav) {
			return true
		}
	}
	return false
}

// addIfChild adds the child navigator to the children if the child.ParentID matches.
func (nav *NavElement) addIfChild(child NavElement) bool {
	if nav.ID == child.ParentID {
		nav.Children = nav.Children.appendSort(child)
		return true
	}
	return false
}
