package main

import (
	"sort"
	"testing"
)

func TestCustomSorting(t *testing.T) {
	var tracks = []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m50s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	}

	sorting := customSort{tracks, []OrderType{Year}, func(x, y *Track, orderBy []OrderType) bool {
		for _, v := range orderBy {
			switch v {
			case Title:
				if x.Title != y.Title {
					return x.Title < y.Title
				}
			case Year:
				if x.Year != y.Year {
					return x.Year < y.Year
				}
			case Length:
				if x.Length != y.Length {
					return x.Length < y.Length
				}
			default:
				return false
			}
		}
		return false
	}}
	sort.Sort(sorting)

	if tracks[0].Year != 1992 {
		t.Fatal()
	}

	sorting.orderBy = []OrderType{Title, Length}
	sort.Sort(sorting)
	if tracks[0].Year != 2012 || tracks[3].Title != "Ready 2 Go" {
		t.Fatal()
	}
}
