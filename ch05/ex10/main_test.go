package main

import "testing"

func TestTopoSort(t *testing.T) {
	n := newTopoSort(prereqs)
	earned := []string{}

	for _, v := range n {
		for item, deps := range prereqs {
			if v == item {
				res := checkDependencies(deps, earned)
				if !res {
					t.Fatalf("You don't have enough units to join class '%v. Earned: %v, Needed: %v", item, earned, deps)
				}
			}
		}

		earned = append(earned, v)
	}
}

func checkDependencies(deps, earned []string) bool {
	for _, dep := range deps {
		found := false
		for _, v := range earned {
			found = found || dep == v
		}
		if !found {
			return false
		}
	}

	return true
}
